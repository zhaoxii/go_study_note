package main

import "fmt"

/**
匿名函数的使用方式：

1、定义匿名函数时直接调用
2、把匿名函数赋值给一个变量，通过该变量来调用匿名函数

*/

func main() {
	fmt.Println("hello world")
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(a)

	b := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(b(10, 20))

}
