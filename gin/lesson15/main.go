package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(ctx *gin.Context) {
		//从请求中读取文件
		fileObj, err := ctx.FormFile("f1") //从请求中获取携带的参数一样的
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取的文件保存在本地(服务器端)
			// dst:=fmt.Sprintf("./%s",fileObj.Filename)
			dst := path.Join("./", fileObj.Filename)
			_ = ctx.SaveUploadedFile(fileObj, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.Run(":9000")
}
