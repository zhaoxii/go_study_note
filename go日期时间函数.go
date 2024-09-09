package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(time.Now())
	fmt.Printf("%T\n", time.Now())
	fmt.Println(time.Now().Year(), int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Printf("%d-%d-%d %d:%d:%d\n", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	timeStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Println(timeStr)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 这个数字是固定的。 必须这样写
	fmt.Println(time.Now().Format("2006/01/02 15:04:05")) // 这个数字是固定的
	fmt.Println(time.Now().Format("2006#01#02 15:04:05")) // 这个数字是固定的
	fmt.Println(time.Now().Format("15:04:05"))            // 数字只写时间就只有时间
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Printf("%T\n", time.Now().Format("2006-01-02"))
	fmt.Println(time.Second, time.Hour, time.Minute, time.Microsecond)
	fmt.Printf("%T, %T, %T, %T\n", time.Second, time.Hour, time.Minute, time.Microsecond)
	time.Sleep(time.Second * 2)
	fmt.Println("=========sleep finish=========")

	fmt.Println(time.Now().Unix(), time.Now().UnixNano(), time.Now().UnixMilli()) // 时间戳
}
