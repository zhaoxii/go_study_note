package main

import "fmt"

/*
string 底层是一个byte数组，因此string也可以进行切片处理



*/

func main() {
	fmt.Println("hello world")

	a := "abcdef"
	slice1 := a[1:4]

	for index, value := range slice1 {
		fmt.Printf("%d\t%c\n", index, value)
	}
	// 字符串不可修改，因为字符串 是不可变的。slice1[0] = "1"（错误的） 要改字符串的话，需要先转为切片，在转为字符串

	slice2 := []byte(a)
	slice2[2] = 'q' // 不能处理中文。一个汉字是3个字节
	a = string(slice2)
	fmt.Println(a)

	slice3 := []rune(a)
	slice3[2] = '赵'
	a = string(slice3)
	fmt.Println(a)

}
