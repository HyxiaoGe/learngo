package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage)
	router.Run(":8083")
}

func startPage(context *gin.Context) {
	var person Person
	if context.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	context.String(http.StatusOK, "Success")
}
