package main

import (
	"fmt"
	"time"
)

// 注意 chan的阻塞机制。 写阻塞，读阻塞

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		i, ok := <-intChan
		fmt.Println("readData:", i, ok)
		if !ok {
			exitChan <- true
			close(exitChan)
			break
		}
		fmt.Println("读到数据", i)
	}
}

func main() {

	intch := make(chan int, 50)
	exitch := make(chan bool, 1)
	go writeData(intch)
	go readData(intch, exitch)

	// 三种等待readData读取完之后再结束的写法：
	//1、
	//for {
	//	val, ok := <-exitch
	//	fmt.Println(val, ok)
	//	if val {
	//		break
	//	}
	//}

	//2、
	//for val := range exitch {
	//	fmt.Println("main", val)
	//}

	// 3.
	val := <-exitch // 从 exitChan 中读取，阻塞等待
	fmt.Println("main", val)

}
