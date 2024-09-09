package main

import (
	"fmt"
)

/*
*
go 不支持 try..except..finally
go 中使用defer， panic， recover来处理

go可以抛出一个panic的异常，然后在defer中通过recover捕获，然后正常处理


*/

func test1() {
	defer func() {
		err := recover() // recover是一个内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("===========")
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	i := num1 / num2
	fmt.Println(i)
}

func main() {

	fmt.Println("aaaaaaa")
	test1()
	fmt.Println("hhhhhhhhh")
}
