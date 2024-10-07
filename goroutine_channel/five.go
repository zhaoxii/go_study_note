package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("=========")
		}
	}()
	var myMap map[int]string
	myMap[0] = "aa"
}

func main() {
	fmt.Println("hello world")
	go hello()
	go test() // 一个goroutine 报错导致整个程序崩溃。在一个goroutine中处理异常，使其不影响其余goroutine
	time.Sleep(time.Second * 10)
	fmt.Println("main end")

}
