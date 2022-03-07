package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//lesson18 中间件
//handlerFunc
func indexHandler(ctx *gin.Context) {
	fmt.Println("index")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//定义一个中间件 HandlerFunc
func m1(c *gin.Context) {
	fmt.Println("m1 in ....")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
	// c.Abort() //阻止调用后续的处理函数
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ....")
	c.Next() //调用后续的处理函数
	fmt.Println("m2 out ...")
	// c.Abort() //阻止调用后续的处理函数
}
func main() {
	r := gin.Default()
	r.Use(m1, m2)
	r.GET("/index1", indexHandler)
	r.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.Run(":9000")
}
