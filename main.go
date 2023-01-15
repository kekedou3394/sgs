package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kekenalog/sgs/middleware"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.New()
	r.Use(middleware.LoggerToFile())

	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{ // c.JSON: 返回 JSON 格式的数据
			"message": "Hello world! I from Gin.",
		})
	})

	// 启动 HTTP 服务，监听 8080 端口
	r.Run(":8080")

}
