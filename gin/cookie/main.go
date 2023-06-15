package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/cookie", func(context *gin.Context) {

		cookie, err := context.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
