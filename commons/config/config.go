package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var GlobalConfig *Config

type Config struct {
	Service  Service  `yaml:"service"`
	DataBase DataBase `yaml:"database"`
	Cache    Cache    `yaml:"cache"`
	Logs     Logs     `yaml:"logs"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type DataBase struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
}

type Cache struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbNum    int    `yaml:"db_num"`
}

type Logs struct {
	AppPath     string `yaml:"app_path"`
	RequestPath string `yaml:"request_path"`
	MaxSize     int    `yaml:"max_size"`
	MaxAge      int    `yaml:"max_age"`
}

// LoadConfig 读取并解析 YAML 配置文件
// 参数：
//   - filePath: 配置文件路径

// MustLoadConfig 加载配置，失败时立即终止程序
func loadConfig(path string) *Config {
	// 1. 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("❌ 配置文件不存在: %s", path)
	}

	// 2. 读取文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("❌ 读取配置文件失败: %v", err)
	}

	// 3. 初始化配置对象
	config := &Config{}

	// 4. 解析YAML内容
	if err = yaml.Unmarshal(data, config); err != nil {
		log.Fatalf("❌ YAML解析失败: %v", err)
	}

	return config
}

func InitConfig(path string) {
	GlobalConfig = loadConfig(path)
}
