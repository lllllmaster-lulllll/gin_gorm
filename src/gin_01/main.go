package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}
func main() {
	r := gin.Default() //返回默认的路由引擎
	r.GET("/hello", sayHello)
	r.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"method": "DELETE",
		})
	})
	//启动服务
	r.Run(":9090")
}
