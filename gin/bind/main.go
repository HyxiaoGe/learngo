package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	router := gin.Default()
	router.POST("/login", func(context *gin.Context) {
		// 可以使用显式绑定声明绑定 multipart form
		// c.shouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定
		var form LoginForm
		// 在这种情况下，将自动选择合适地绑定
		if context.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run()
}
