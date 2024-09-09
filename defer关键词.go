package main

import "fmt"

/**

在函数中，通常需要创建资源，比如，数据库连接，文件句柄，锁等。为了在函数执行完及时的释放资源，go提供 defer 延时机制

*/

func sum1(x int, y int) int {
	// 当执行到defer时，会将 defer后的语句压入独立的栈中， 暂时不执行。当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈执行。
	defer fmt.Println("x=", x)
	defer fmt.Println(y)
	// defer 将语句放入栈中时，也会将相关的值拷贝同时入栈。
	x++
	y++
	res := x + y
	fmt.Println("res=", res)
	return res
}

func main() {
	fmt.Println("hello world")
	fmt.Println(sum1(10, 20))

}
