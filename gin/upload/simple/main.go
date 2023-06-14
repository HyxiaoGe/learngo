package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()
	// 为 multipart form 设置较低的内存限制（默认是 32MiB）
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/upload", func(context *gin.Context) {
		// 单文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		context.SaveUploadedFile(file, dst)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})

	router.Run()

}
