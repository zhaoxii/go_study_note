package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//var chan1 chan int // 默认是可读可写

	//var chan2 chan<- int  // 只写
	//var chan2 <-chan int  // 只读
	//	只读和只写是对chan的限制，不是新的类型

	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		stringChan <- "hello" + strconv.Itoa(i)
	}

	intChan := make(chan int, 5)
	for i := 0; i < 5; i++ {
		intChan <- i
	}

	// select 可以解决从管道取数据的阻塞问题。如果多个 case 同时准备好，Go 会随机选择一个执行。
	for {
		select {
		case v := <-stringChan: // 这里如果管道一直没有关闭，也不会一直阻塞导致死锁。会自动的到下一个case匹配
			fmt.Println("v", v)
		case v := <-intChan:
			fmt.Println("v", v)
		default:
			fmt.Println("default")
			return
		}
		time.Sleep(time.Second)
	}

}
