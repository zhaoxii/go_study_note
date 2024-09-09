package main

import "fmt"

func main() {
	count := 0
	var arr [20]int
	for i := 0; i <= 100; i++ {
		if i%9 == 0 {
			arr[count] = i
			count++
		}
	}
	fmt.Println(arr)
	fmt.Println(count)
}
