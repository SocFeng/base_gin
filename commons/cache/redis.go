package cache

import (
	"base_gin/commons/config"
	"base_gin/commons/logs"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var GlobalRedis *redis.Client

func createRedis() {

}

// InitRedis 初始化 Redis
func InitRedis() {
	GlobalRedis = redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379", // Redis 服务器地址
		Addr:     fmt.Sprintf("%s:%d", config.GlobalConfig.Cache.Host, config.GlobalConfig.Cache.Port), // Redis 服务器地址
		Password: config.GlobalConfig.Cache.Password,                                                   // 没有密码时留空
		DB:       config.GlobalConfig.Cache.DbNum,                                                      // 默认 DB
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := GlobalRedis.Ping(ctx).Result()
	if err != nil {
		logs.AppError(" ❌ 无法连接到 Redis: %v", err)
	}

	logs.AppInfo(" ✅ Redis 连接成功")
}
