package middleware

import (
	"base_gin/commons/logs"
	"github.com/gin-gonic/gin"
)

// RequestLogger 记录请求日志的中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取的数据
		requestData, _ := GetRequestDataBytes(c)

		// 处理请求
		c.Next()
		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		logs.RequestInfo(map[string]interface{}{
			"client_ip":   clientIP,
			"method":      method,
			"path":        path,
			"requestData": string(requestData),
		})
	}
}
