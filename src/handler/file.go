package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 上传单个文件
func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("上传文件出错")
	}
	fmt.Println("fileName : ", file.Filename)

	err = c.SaveUploadedFile(file, file.Filename)
	fmt.Println("是否错误：", err)
}

/**
上传多个文件:注意文件覆盖问题
*/
func MultiFileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	for key, file := range files {
		fmt.Println("key = ", key, "filename", file.Filename)
		c.SaveUploadedFile(file, "file/"+file.Filename)
	}

}
