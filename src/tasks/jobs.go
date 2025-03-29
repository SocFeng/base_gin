package tasks

import (
	"base_gin/commons/logs"
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

// InitCron 初始化定时任务
func InitCron() {
	Cron = cron.New(cron.WithSeconds())

	// 添加定时任务
	_, err := Cron.AddFunc("*/10 * * * * *", DoNumber)
	if err != nil {
		logs.AppError("DoNumber job is err:", err)
	}

	_, err = Cron.AddFunc("0 0 * * * *", MidJob)
	if err != nil {
		logs.AppError("DoNumber job is err:", "每天午夜执行的任务", err)
	}

	// 启动定时任务
	Cron.Start()
}

// StopCron 停止定时任务
func StopCron() {
	Cron.Stop()
}
