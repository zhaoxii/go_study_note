package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义一个中间件m1（函数，参数是*gin.Context类型的）
func m1(c *gin.Context) {
	fmt.Println("m1...")
	start := time.Now()
	c.Set("name", "zx")
	c.Next() //调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	fmt.Println("花费多长时间: ", time.Since(start))
	fmt.Println("m1...out")
}

func m2(c *gin.Context) {
	fmt.Println("m2...")
	c.Abort()
	fmt.Println("m2...out")
}

func main() {
	r := gin.Default()

	//r.Use(m1) // 全局注册中间件
	r.Use(m1, m2)

	//r.GET("/", m1, func(c *gin.Context) {
	//	fmt.Println("hello gin")
	//	c.JSON(200, gin.H{"hello": "world"})
	//}) // 先执行m1， 在执行后面的函数。 m1就相当于一个中间件。

	r.GET("/", func(c *gin.Context) {
		fmt.Println("hello gin")
		name, ok := c.Get("name")
		if !ok {
			name = "xxx"
		}
		c.JSON(200, gin.H{"hello": "world", "name": name})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
