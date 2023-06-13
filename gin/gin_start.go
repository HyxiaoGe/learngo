package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// "github.com/gin-gonic/gin"
func main() {

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 默认在 0.0.0.0:8080 上启动服务并监听
	router.Run()

}
