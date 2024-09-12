package main

import "fmt"

/*
结构体方法

golang 中的 方法是作用在指定的数据类型上的（和指定的数据类型绑定），因此自定义类型，都可以有方法。
而不仅仅是struct。int，float32等都可以有方法

方法的声明和调用

	type A struct{
		Num int
	}

	func (a A)test(参数列表) (返回值类型) {
		fmt.Println("test")
	}

表示A结构体有个方法 test。 （a A）表示test方法是和A结构体绑定的。 a表示是调用方法的对象实例副本。在方法内改数据不会影响到外面


方法的访问范围和函数一样，方法名首字母小写，只能在本包使用，大写，可以在其他包使用

如果一个变量实现了String()方法。那么fmt.Println默认会调用这个变量的String()输出。

*/

type A struct {
	Num int
}

type Person1 struct {
	name string
}

type integer int

func (i integer) print() {
	fmt.Printf("i=%d\n", i)
}

func (i integer) String() string {
	return "hhhhhhhhhhhh"
}

func (a A) test() {
	fmt.Println("test")
	fmt.Println(a.Num)
	a.Num = 10000 // 不会影响到外面。值传递
}

func (p Person1) speak() {
	fmt.Println(p.name, "speak")
}

func (p Person1) js() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.name, res)

}

func (p Person1) jsn(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(p.name, res)
}

func (p Person1) jsnreturn(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func (p Person1) String() string {
	fmt.Println("zxzs")
	return "zxzs"
}

func main() {

	fmt.Println("hello world")
	var a = A{Num: 100}
	a.test()
	fmt.Println(a)

	p := Person1{name: "xxx"}
	p.speak()
	p.js()
	p.jsn(10)
	res := p.jsnreturn(10)
	fmt.Println(res)
	fmt.Println(p)
	fmt.Println("==================================")

	var i integer = 10
	i.print()
	fmt.Println(i)

}
