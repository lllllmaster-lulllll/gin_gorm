package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发送请求携带的 query string 参数
		name := c.Query("query") // 通过 query 获取请求中的参数
		age := c.Query("age")
		// name := c.DefaultQuery("query", "kn")//取不到就用指定的默认值
		// name, ok := c.GetQuery("query") //取不到第二个参数就返回 false
		// if !ok {
		// 	fmt.Println("query 未获取到")
		// 	name = "somebody"
		// }
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9000")
}
