package main

import "fmt"

/**

闭包就是一个函数和与其相关的引用环境组合的一个整体。

*/

func AddUpper() func(int) int {
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}

/*
*
addupper返回一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包。
可以这样理解： 闭包是类， 函数是操作，n是字段。 函数和它使用到的n构成闭包。
当我们反复调用函数时，因为n只初始化一次，因此每次调用就累加。
我们要搞清楚闭包的关键，就是要分析出返回的函数使用到哪些变量，因为函数和它引用的变量共同构成闭包。
*/
func main() {

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(2))

}
