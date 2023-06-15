package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// 使用现有的基础请求对象解析查询字符串参数
	// 示例 URL: /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "guest")
		lastname := context.Query("lastname") // context.URL.Query("lastname") 的一种快捷方式

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run()
}
