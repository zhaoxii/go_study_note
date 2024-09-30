package main

import "fmt"

/*

接口(interface)

定义一个接口, 声明了两个没有实现的方法
type Usb interface{
	Start()
	Stop()
}

interface类型可以定义一组方法。但是不需要实现。并且interface不能包含任何变量。到某个自定义类型(比如结构体Phone)要使用的时候，把这些方法写出来

基本语法：
type 接口名称 interface{
	method1(参数列表) 返回值列表
	method2(参数列表) 返回值列表
	...
}

实现接口：定义结构体，并且实现该接口所有方法。（没有实现所有方法不是实现接口）

接口里的所有方法都是没有实现的方法。

接口体现了程序设计的多态和高内聚低耦合的思想

golang中的接口，不需要显式的实现(implement)。只要有一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。
因此，golang中没有implement这样的关键字

*/

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

// 让phone实现usb接口的方法

func (p Phone) Start() {
	fmt.Println("phone start")
}

func (p Phone) Stop() {
	fmt.Println("phone stop")
}

type Camera struct {
}

// 让Camera实现usb接口的方法

func (c Camera) Start() {
	fmt.Println("c Camera start")
}

func (c Camera) Stop() {
	fmt.Println("c Camera stop")
}

type Computer struct {
}

func (c Computer) Work(usb Usb) {
	//方法接受一个usb接口类型的参数，只要是实现了usb接口(实现usb接口就是指实现了usb接口声明的所有方法)
	fmt.Println("work computer")
	// 通过usb接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	fmt.Println("hello world")
	computer1 := Computer{}
	phone1 := Phone{}
	computer1.Work(phone1)
	camera1 := Camera{}
	computer1.Work(camera1)
}
