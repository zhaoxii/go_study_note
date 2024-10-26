package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/three", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")

		fmt.Println(name, age)

		c.JSON(200, gin.H{"name": name, "age": age})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
