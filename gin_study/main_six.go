package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Any 支持所有请求方法
	r.Any("/six", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "six"})
	})

	// 当没有访问的路由时，执行的
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "zz"})
	})

	userGroup := r.Group("/user") // 类似fastapi router, flask blueprint

	userGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "user get"})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
