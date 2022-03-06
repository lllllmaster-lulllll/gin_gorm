package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx", "./statics")
	//解析模板之前设置函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP 请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "lllllmaster-lulllll.github.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		//HTTP 请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	//返回从网上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":9000")
}
