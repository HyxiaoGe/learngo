package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	router := gin.Default()
	router.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})
	router.Run()

}
