package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hey", "status": http.StatusOK,
		})
	})

	router.GET("/moreJSON", func(context *gin.Context) {

		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		context.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.Run()

}
