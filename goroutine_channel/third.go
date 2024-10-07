package main

import (
	"fmt"
	"runtime"
	"time"
)

func addData(intChan chan int) {
	for i := 1; i <= 10000; i++ {
		intChan <- i
		fmt.Printf("goroutine addData:%d\n", i)
	}
	close(intChan)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false // 小于等于1的数不是素数
	}
	if num <= 3 {
		return true // 2和3是素数
	}
	if num%2 == 0 || num%3 == 0 {
		return false // 排除能被2和3整除的数
	}
	for i := 5; float64(i*i) <= float64(num); i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func primeAdd(intChan chan int, primeChan chan int, flagChan chan bool, flag int) {

	for {
		number, ok := <-intChan
		fmt.Println("primeAdd, flag", flag, number, ok)
		if ok && isPrime(number) {
			primeChan <- number
		}
		if !ok {
			flagChan <- true
			break
		}
	}

}

func main() {
	fmt.Println("hello world")
	runtime.GOMAXPROCS(runtime.NumCPU())
	curTime := time.Now().UnixMilli()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	flagChan := make(chan bool, 4)

	go addData(intChan)
	go primeAdd(intChan, primeChan, flagChan, 1)
	go primeAdd(intChan, primeChan, flagChan, 2)
	go primeAdd(intChan, primeChan, flagChan, 3)
	go primeAdd(intChan, primeChan, flagChan, 4)

	for {
		fmt.Println("----------------")
		if len(flagChan) == 4 {
			fmt.Println("break-------------")
			break
		}
	}
	close(primeChan)
	//fmt.Println("素数个数: ", len(primeChan))
	//for value := range primeChan {
	//	fmt.Printf("%v, ", value)
	//}
	fmt.Println(time.Now().UnixMilli() - curTime)
	fmt.Println(runtime.NumCPU())

}
