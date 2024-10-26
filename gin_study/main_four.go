package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/four/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")

		fmt.Println("name:", name, "age:", age)
		c.JSON(200, gin.H{"name": name, "age": age})

	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
