package dataBase

import (
	"base_gin/commons/config"
	"base_gin/commons/logs"
	"base_gin/commons/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GlobalGormDB *gorm.DB

func createDsnString() string {
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.DataBase.User,
		config.GlobalConfig.DataBase.Password,
		config.GlobalConfig.DataBase.Host,
		config.GlobalConfig.DataBase.Port,
		config.GlobalConfig.DataBase.DbName)
	return dsnString

}

// InitDB 初始化数据库连接
func InitGormDB() {
	// DSN 格式：用户名:密码@tcp(数据库地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := createDsnString()
	var err error
	GlobalGormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.AppFatal(" ❌ 数据库连接失败:", err)
	}
	logs.AppInfo(" ✅ 数据库连接成功！")

	// 可选：设置连接池配置
	sqlDB, err := GlobalGormDB.DB()
	if err != nil {
		logs.AppFatal("获取数据库实例失败:", err)
	}
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大打开连接数
	sqlDB.SetConnMaxLifetime(1 * time.Hour) // 连接最大存活时间
}

// MigrateDB 自动迁移模型
func MigrateDB() {
	// 自动创建或更新数据库中的表结构
	err := GlobalGormDB.AutoMigrate(&models.User{})
	if err != nil {
		logs.AppFatal("❌ 数据库迁移失败！")
	}
}
