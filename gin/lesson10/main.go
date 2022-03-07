package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("json", func(c *gin.Context) {
		// 方法 1: 使用 map
		// data := map[string]interface{}{
		// 	"name":    "小王子",
		// 	"message": "hello,world",
		// 	"age":     18,
		// }
		data := gin.H{
			"name":    "小王子",
			"message": "hello,world",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})
	//方法 2 使用结构体
	type msg struct {
		Name    string
		Message string
		Age     int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "小王子",
			Message: "hello,GOLANG",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9000")
}
