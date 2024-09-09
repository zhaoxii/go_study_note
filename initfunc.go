package main

import "fmt"

/**
每个源文件都可以有一个init函数，该函数会在main函数执行前被调用

通常可以在init函数中完成初始化的工作

如果一个文件中同时包含全局变量，init函数和main函数，则执行的流程是： 全局变量定义 -> init函数 -> main函数


*/

var age = test()

func test() int {
	fmt.Println("hello test()")
	return 90
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println(age)
	fmt.Println("hello world")
}
