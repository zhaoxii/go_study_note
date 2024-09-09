package main

import "fmt"

/**
函数
func 函数名(形参)（返回值列表）{
	函数体
	return 返回值
}



*/

func sum(x int, y int) (int, int, int) {
	return x + y, x, y
}

func multe(params ...int) {
	fmt.Println(params)
	fmt.Println(params[0], params[1], params[2])
	fmt.Printf("%T\n", params)
}

func main() {

	fmt.Println("Hello World")
	sumRes, x, y := sum(1, 2)
	fmt.Println(x, y, sumRes)
	multe(1, 2, 3)
}
