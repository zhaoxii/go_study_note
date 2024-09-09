package main

import "fmt"

func main() {
	// 一般不主张使用goto
	i := 0
	for {
		if i == 2 {
			goto xxx
		}
		i++
		fmt.Println("eeee")

	}

xxx: // 标签的定义只要符合 标识符就可以
	fmt.Println("xxx")

}
