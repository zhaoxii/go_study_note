package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student2 struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/five", func(c *gin.Context) {

		var s Student2
		err := c.ShouldBind(&s) // Query参数绑定。 注意传地址
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Println(s)
			c.JSON(http.StatusOK, gin.H{"name": s.Name, "age": s.Age})
		}

	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
