package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		context.Set("example", "12345")
		// 请求前
		context.Next()
		// 请求后
		latency := time.Since(t)
		log.Print("latency: ", latency)

		// 获取发送的 status
		status := context.Writer.Status()
		log.Print("status: ", status)
	}
}

func main() {

	router := gin.New()
	router.Use(Logger())

	router.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)

		// 打印：”12345“
		log.Println("example", example)
	})

	router.Run()
}
