package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// u := UserInfo{
		// 	UserName: username,
		// 	Password: password,
		// }
		var u UserInfo //声明一个 UserInfo 结构体
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.POST("/form", func(ctx *gin.Context) {
		u := UserInfo{}
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.POST("/json", func(ctx *gin.Context) {
		u := UserInfo{}
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.Run(":9000")
}
