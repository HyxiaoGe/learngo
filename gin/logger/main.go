package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	router.Run(":8083")
}
