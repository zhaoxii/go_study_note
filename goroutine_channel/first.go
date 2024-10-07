package main

import (
	"fmt"
	"runtime"
	"time"
)

func goroutineTest() {
	fmt.Println("goroutine_test")
	time.Sleep(time.Second * 3)
	fmt.Println("goroutine_test2")
}

func main() {
	fmt.Println("hello world")
	go goroutineTest()
	fmt.Println("hello world2")
	time.Sleep(time.Second * 2)
	num := runtime.NumCPU()
	fmt.Println("num:", num)

	var mychan chan int
	mychan = make(chan int, 10)
	for i := 0; i < num; i++ {
		mychan <- i + 1
	}
	fmt.Println(mychan)
	fmt.Println(len(mychan), cap(mychan))
	num2 := <-mychan
	fmt.Println(num2)
	fmt.Println(len(mychan), cap(mychan))
	<-mychan
	fmt.Println(len(mychan), cap(mychan))
	close(mychan)
	fmt.Println(len(mychan), cap(mychan))
	fmt.Println(<-mychan)
	fmt.Println(len(mychan), cap(mychan))
	//mychan <- 33  // 关闭之后不能写入panic: send on closed channel
	fmt.Println(len(mychan), cap(mychan))

	//	 遍历chan
	for val := range mychan {
		fmt.Println(val)
		fmt.Println("============================")
	}

	mychan2 := make(chan int, 10)
	for i := 0; i < cap(mychan2); i++ {
		mychan2 <- i + 1
	}

	// 需要先关闭chan再遍历，否则会报错
	for val := range mychan2 {
		fmt.Println(val)
		fmt.Println("+++++++++++")
	}

}
