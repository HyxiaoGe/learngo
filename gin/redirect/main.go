package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	//router.GET("/redirect", func(context *gin.Context) {
	//	context.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	//})
	//router.POST("/test", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/foo")
	//})

	router.GET("/test", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		router.HandleContext(context)
	})
	router.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.Run()
}
