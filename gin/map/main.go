package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/post", func(context *gin.Context) {

		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	router.Run()
}
