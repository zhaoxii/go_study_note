package main

import "fmt"

/*
*

len

new  用来分配内存。 主要用来分配值类型, 基本类型。分配内存并返回指向该类型的指针

make 用来分配内存。 主要是分配引用类型。 chan， slice，map
*/
func main() {

	i := new(int)
	fmt.Println(i)
	fmt.Println(*i)
	fmt.Printf("%T\n", i)

	*i = 1000
	fmt.Println(*i)
}
