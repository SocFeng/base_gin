package main

import (
	"base_gin/commons/cache"
	"base_gin/commons/config"
	"base_gin/commons/dataBase"
	"base_gin/commons/logs"
	"base_gin/commons/middleware"
	"base_gin/src/api"
	"base_gin/src/tasks"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 初始化配置文件
	config.InitConfig("config.yaml")
	fmt.Println(config.GlobalConfig.Service.Host)
	fmt.Println(config.GlobalConfig.Service.Port)

	// 配置日志
	logs.InitLog()
	logs.AppInfo("👌👌👌👌 Hello   !")

	// 初始化数据库
	dataBase.InitDB()
	defer func(GlobalDB *sql.DB) {
		err := GlobalDB.Close()
		if err != nil {

		}
	}(dataBase.GlobalDB)
	// 初始化gorm 并且数据库迁移
	dataBase.InitGormDB()
	dataBase.MigrateDB()

	// 初始化缓存数据库redis
	cache.InitRedis()
	defer func(GlobalRedis *redis.Client) {
		err := GlobalRedis.Close()
		if err != nil {

		}
	}(cache.GlobalRedis)

	tasks.InitCron()
	defer tasks.StopCron()

	serviceString := fmt.Sprintf("%s:%d", config.GlobalConfig.Service.Host, config.GlobalConfig.Service.Port)

	r := gin.Default()
	r.Use(middleware.MergeRequestData(), middleware.RequestLogger())
	api.InitRoutes(r)
	if err := r.Run(serviceString); err != nil {
		logs.AppFatal("【ERROR】  ❌  服务器启动失败！", err)
	} else {
		logs.AppInfo("【INFO】  ✅  服务器启动成功！", err)
	}

}
