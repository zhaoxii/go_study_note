package main

import "fmt"

func getSum(x, y int) int {
	return x + y
}

func myFun(funicular func(int, int) int, x int, y int) int {
	return funicular(x, y)
}

func main() {
	res := myFun(getSum, 10, 20)
	fmt.Println(res)
}
