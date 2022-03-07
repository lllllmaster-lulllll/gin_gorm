package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//form表单提交与获取
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//login post
	r.POST("/login", func(c *gin.Context) {
		// 1. PostForm获取 form 表单提交的数据
		// username := c.PostForm("username")
		// password := c.PostForm("password")
		//2.DefaultPostForm
		// username := c.DefaultPostForm("username", "somebody")
		// password := c.DefaultPostForm("password", "***")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "pw"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9000")
}
