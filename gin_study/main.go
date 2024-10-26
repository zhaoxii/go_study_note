package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		//data := map[string]interface{}{
		//	"name": "sss",
		//}

		//c.JSON(200, data)
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	r.GET("/user", func(c *gin.Context) {
		u1 := Student{
			Name: "zzz",
			Age:  18,
		}
		c.JSON(http.StatusOK, u1)
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
