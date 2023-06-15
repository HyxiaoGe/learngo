package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {

	router := gin.Default()

	// 绑定 JSON({"user": "admin", "password": "123456"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if json.User != "admin" || json.Password != "123456" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logger in"})
	})

	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginForm", func(context *gin.Context) {
		var form Login
		// 根据 Content-Type Header 推断使用哪个绑定器
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "admin" || form.Password != "123456" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logger in"})
	})

	router.Run()
}
