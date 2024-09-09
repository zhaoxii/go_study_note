package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 0
	for {

		// 生成一个 0 到 99 之间的随机整数
		randomNum := rand.Intn(100)
		if randomNum == 99 {
			fmt.Println("generate count %d", count+1)
			break
		}
		fmt.Println("随机整数:", randomNum)
		count += 1

	}

lable: // 外层
	for i := 0; i < 4; i++ {
	xxx: //设置一个标签
		for j := 0; j < 10; j++ {
			if j == 1 {
				break lable
			}
			if j == 2 {
				break xxx
			}
			fmt.Printf("j=%d ", j)
		}
	}

	// 多层循环中， 默认跳出一层，但可以指定标签，跳出到指定处。  continue 也可以这样使用
}
