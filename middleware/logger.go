package middleware

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kekenalog/sgs/config"
	"github.com/sirupsen/logrus"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	var src *os.File
	// 	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		src, err = os.Create(fileName)
	} else {
		src, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = io.MultiWriter(src, os.Stdout)

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// var AppLog *logrus.Logger
// var WebLog *logrus.Logger

// func Setup() {
// 	initAppLog()
// 	initWebLog()
// }

// /*
// *

// 	初始化AppLog
// */
// func initAppLog() {
// 	logFileName := "app.log"
// 	AppLog = initLog(logFileName)
// }

// /*
// *

// 	初始化WebLog
// */
// func initWebLog() {
// 	logFileName := "web.log"
// 	WebLog = initLog(logFileName)
// }

// /*
// *

// 	初始化日志句柄
// */
// func initLog(logFileName string) *logrus.Logger {
// 	log := logrus.New()
// 	log.Formatter = &logrus.JSONFormatter{
// 		TimestampFormat: "2006-01-02 15:04:05",
// 	}
// 	logPath := "./logs/"
// 	logName := logPath + logFileName
// 	var f *os.File
// 	var err error
// 	//判断日志文件是否存在，不存在则创建，否则就直接打开
// 	if _, err := os.Stat(logName); os.IsNotExist(err) {
// 		f, err = os.Create(logName)
// 	} else {
// 		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	}

// 	if err != nil {
// 		fmt.Println("open log file failed")
// 	}

// 	log.Out = f
// 	log.Level = logrus.InfoLevel
// 	return log
// }

// /*
// *

// 	Gin中间件函数，记录请求日志
// */
// func LoggerToFile() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 开始时间
// 		startTime := time.Now()

// 		// 处理请求
// 		c.Next()

// 		// 结束时间
// 		endTime := time.Now()

// 		// 执行时间
// 		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))

// 		// 请求方式
// 		reqMethod := c.Request.Method

// 		// 请求路由
// 		reqUri := c.Request.RequestURI

// 		// 状态码
// 		statusCode := c.Writer.Status()

// 		// 请求IP
// 		clientIP := c.ClientIP()

// 		//日志格式
// 		WebLog.WithFields(logrus.Fields{
// 			"http_status": statusCode,
// 			"total_time":  latencyTime,
// 			"ip":          clientIP,
// 			"method":      reqMethod,
// 			"uri":         reqUri,
// 		}).Info("access")
// 	}
// }
