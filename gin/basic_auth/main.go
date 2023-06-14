package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {

	router := gin.Default()
	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 ”localhost:8080/admin/secrets“
	authorized.GET("/secrets", func(context *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := context.MustGet(gin.AuthUserKey).(string)
		if secrets, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secrets})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRETS :("})
		}
	})

	router.Run(":8083")
}
