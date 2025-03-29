package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Config 中间件配置选项
type Config struct {
	MaxMemory        int64 // 最大内存限制（用于解析 multipart/form-data，默认 32MB）
	EnableFile       bool  // 是否处理文件上传字段
	AllowMultiValues bool  // 是否保留多值字段（如 query=1&query=2 存储为 []string）
}

// DefaultConfig 默认配置
var DefaultConfig = Config{
	MaxMemory:        32 << 20, // 32MB
	EnableFile:       true,     // 允许传递文件
	AllowMultiValues: true,     // 允许传递多个文件
}

// MergeRequestData 中间件（分类存储 query、json、form 数据）
func MergeRequestData(config ...Config) gin.HandlerFunc {
	cfg := DefaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(c *gin.Context) {
		// 统一存储 requestData，并分类存储
		requestData := map[string]interface{}{
			"query": map[string]interface{}{},
			"json":  map[string]interface{}{},
			"form":  map[string]interface{}{},
			"files": map[string]interface{}{},
		}

		// 1. 解析 Query 参数
		queryParams := requestData["query"].(map[string]interface{})
		query := c.Request.URL.Query()
		for key, values := range query {
			if cfg.AllowMultiValues && len(values) > 1 {
				queryParams[key] = values // 存储多值
			} else {
				queryParams[key] = values[0] // 存储单值
			}
		}

		// 2. 解析 Body 数据
		contentType := c.ContentType()
		switch contentType {
		case "application/json":
			if err := parseJSON(c, requestData["json"].(map[string]interface{})); err != nil {
				abortWithError(c, http.StatusBadRequest, "JSON 解析失败", err)
				return
			}

		case "application/x-www-form-urlencoded", "multipart/form-data":
			if err := parseForm(c, requestData["form"].(map[string]interface{}), requestData["files"].(map[string]interface{}), cfg); err != nil {
				abortWithError(c, http.StatusBadRequest, "表单解析失败", err)
				return
			}
		}

		// 3. 注入上下文
		c.Set("requestData", requestData)
		c.Next()
	}
}

// parseJSON 解析 JSON 数据并合并到 requestData
func parseJSON(c *gin.Context, data map[string]interface{}) error {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		return err
	}
	mergeMaps(data, jsonData)
	return nil
}

// parseForm 解析表单数据（支持文件上传）
func parseForm(c *gin.Context, formData map[string]interface{}, fileData map[string]interface{}, cfg Config) error {
	if contentType := c.ContentType(); contentType == "multipart/form-data" {
		if err := c.Request.ParseMultipartForm(cfg.MaxMemory); err != nil {
			return err
		}
	} else {
		if err := c.Request.ParseForm(); err != nil {
			return err
		}
	}

	// 处理表单字段
	form := c.Request.PostForm
	for key, values := range form {
		if cfg.AllowMultiValues && len(values) > 1 {
			formData[key] = values
		} else {
			formData[key] = values[0]
		}
	}

	// 处理文件上传
	if cfg.EnableFile && c.Request.MultipartForm != nil {
		files := c.Request.MultipartForm.File
		for field, headers := range files {
			fileInfo := make([]map[string]interface{}, len(headers))
			for i, h := range headers {
				fileInfo[i] = map[string]interface{}{
					"filename": h.Filename,
					"size":     h.Size,
					"header":   h.Header,
				}
			}
			fileData[field] = fileInfo
		}
	}

	return nil
}

// mergeMaps 深度合并两个 map
func mergeMaps(dest, src map[string]interface{}) {
	for key, srcVal := range src {
		if destVal, exists := dest[key]; exists {
			if srcMap, ok := srcVal.(map[string]interface{}); ok {
				if destMap, ok := destVal.(map[string]interface{}); ok {
					mergeMaps(destMap, srcMap)
					continue
				}
			}
		}
		dest[key] = srcVal
	}
}

// abortWithError 中断请求并返回错误
func abortWithError(c *gin.Context, code int, message string, err error) {
	log.Printf("[Middleware Error] %s: %v", message, err)
	c.AbortWithStatusJSON(code, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}

// GetRequestData 从上下文中获取合并后的数据
func GetRequestData(c *gin.Context) map[string]interface{} {
	data, exists := c.Get("requestData")
	if !exists {
		return nil
	}
	return data.(map[string]interface{})
}

// GetString 从请求数据中安全获取字符串
func GetString(data map[string]interface{}, key string) string {
	val, ok := data[key]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v", val)
}

// GetSlice 获取多值字段切片
func GetSlice(data map[string]interface{}, key string) []string {
	val, ok := data[key]
	if !ok {
		return nil
	}

	if slice, ok := val.([]string); ok {
		return slice
	}

	return []string{fmt.Sprintf("%v", val)}
}

// GetRequestDataBytes 将合并后的请求数据序列化为 JSON 格式的 []byte
func GetRequestDataBytes(c *gin.Context) ([]byte, error) {
	data := GetRequestData(c) // 从上下文中获取 map[string]interface{}
	return json.Marshal(data)
}

// GetRequestDataBytesIndent 将数据序列化为格式化的 JSON（便于调试）
func GetRequestDataBytesIndent(c *gin.Context) ([]byte, error) {
	data := GetRequestData(c)
	return json.MarshalIndent(data, "", "  ")
}
