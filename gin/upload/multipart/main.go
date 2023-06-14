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
		// multipart form
		form, _ := context.MultipartForm()
		files := form.File["files[]"]

		for _, file := range files {
			log.Println(file.Filename)

			dst := "./" + file.Filename
			// 上传文件至指定目录
			context.SaveUploadedFile(file, dst)
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})

	router.Run()

}
