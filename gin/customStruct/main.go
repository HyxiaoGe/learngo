package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		Field string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(context *gin.Context) {
	var b StructB
	context.Bind(&b)
	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(context *gin.Context) {
	var c StructC
	context.Bind(&c)
	context.JSON(http.StatusOK, gin.H{
		"a": c.NestedStructPointer,
		"c": c.FieldC,
	})
}

func GetDataD(context *gin.Context) {
	var d StructD
	context.Bind(&d)
	context.JSON(http.StatusOK, gin.H{
		"a": d.NestedAnonyStruct,
		"d": d.FieldD,
	})
}

func main() {

	router := gin.Default()
	router.GET("/getb", GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)

	router.Run()
}
