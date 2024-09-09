package main

import "fmt"

// go 没有 while 和 do while。 可以使用for 循环实现这样的效果
func main() {
	var i = 0
	for {
		if i > 2 {
			break
		}
		fmt.Println("xxx")
		i++
	}
	fmt.Println(i)

}
