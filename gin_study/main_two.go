package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "12")
		skill := c.QueryArray("skill")
		fmt.Println(age, name, skill)
		c.JSON(200, gin.H{"name": name, "age": age, "skill": skill})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
