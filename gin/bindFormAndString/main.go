package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {

	router := gin.Default()
	router.GET("/testing", startPage)
	router.Run()

}

func startPage(context *gin.Context) {
	var person Person
	// 如果是 GET 请求，只能使用 Form 绑定引擎
	// 如果是 POST 请求，首先检查 content-type 是否为 JSON 或者 XML，然后再使用 Form（form-data）
	if context.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	context.JSON(http.StatusOK, "Success")
}
