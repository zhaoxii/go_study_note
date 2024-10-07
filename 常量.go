package main

import "fmt"

/*
常量使用const修改
常量在定义的时候，必须初始化
常量不能修改
常量只能修饰bool， 数值类型(int，float系列)，string类型
语法：const identifier [type] = value

仍然是通过首字母大小写来控制包的访问

*/

const name = "aaa"

const (
	a1 = iota
	a2
	a3
	a4
	a5
	a6
	a7
)

func main() {
	fmt.Println("hello world")
	fmt.Println(a1, a2, a3, a4, a5, a6, a7)
}
