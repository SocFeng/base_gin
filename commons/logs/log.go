package logs

import (
	"base_gin/commons/config"
	"fmt"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	AppLogger     *logrus.Logger
	RequestLogger *logrus.Logger
)

func createLogger(fileName string) *logrus.Logger {
	logger := logrus.New()
	// 设置日志输出到文件和控制台
	logFile := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    config.GlobalConfig.Logs.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     config.GlobalConfig.Logs.MaxAge, // days
		Compress:   true,
	}

	// 设置日志输出到多个目的地
	mw := io.MultiWriter(logFile)
	logger.SetOutput(mw)
	// 设置日志格式
	logger.SetFormatter(&customFormatter{})
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 启用调用者信息
	logger.SetReportCaller(true)

	return logger
}

// InitLog 初始化日志
func InitLog() {
	AppLogger = createLogger(config.GlobalConfig.Logs.AppPath)
	RequestLogger = createLogger(config.GlobalConfig.Logs.RequestPath)
}

// customFormatter 自定义日志格式
type customFormatter struct{}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006/01/02 15:04:05")
	level := strings.ToUpper(entry.Level.String()) // 将日志级别转换为大写
	msg := entry.Message

	log := fmt.Sprintf("[%s] %s --> %s\n", level, timestamp, msg)

	return []byte(log), nil
}

// Debug 记录调试日志
func AppDebug(args ...interface{}) {
	AppLogger.Debug(args...)
}

// Info 记录信息日志
func AppInfo(args ...interface{}) {
	AppLogger.Info(args...)
}

// Warn 记录警告日志
func AppWarn(args ...interface{}) {
	AppLogger.Warn(args...)
}

// Error 记录错误日志
func AppError(args ...interface{}) {
	AppLogger.Error(args...)
}

// Fatal 记录致命错误日志并退出程序
func AppFatal(args ...interface{}) {
	AppLogger.Fatal(args...)
}

// Debug 记录调试日志   ======================== Request =======================
func RequestDebug(args ...interface{}) {
	RequestLogger.Debug(args...)
}

// Info 记录信息日志
func RequestInfo(args ...interface{}) {
	RequestLogger.Info(args...)
}

// Warn 记录警告日志
func RequestWarn(args ...interface{}) {
	RequestLogger.Warn(args...)
}

// Error 记录错误日志
func RequestError(args ...interface{}) {
	RequestLogger.Error(args...)
}

// Fatal 记录致命错误日志并退出程序
func RequestFatal(args ...interface{}) {
	RequestLogger.Fatal(args...)
}
