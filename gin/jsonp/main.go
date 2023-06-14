package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.GET("/JSONP", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSON(http.StatusOK, data)
	})

	router.Run()
}
