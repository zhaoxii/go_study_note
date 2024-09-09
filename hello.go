package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("hello!!")
	const name string = "zhaoxi"
	fmt.Println(name)
	const age = 1222
	fmt.Println(age)
	const a = iota
	fmt.Println(a)

	sex := "男"
	fmt.Println(sex)

	var name1 = "lisi"
	i := len(name1)
	fmt.Println(i)
	func1()

	var sli_1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)
	intOs := make([]int, 5)
	fmt.Println(intOs)

	var p1 Person
	p1.Name = "Tom"
	p1.Age = 30
	fmt.Println("p1 =", p1)
	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2 =", p2)
	//匿名结构体
	p4 := struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 33}
	fmt.Println("p4 =", p4)
}

func func1() {
	var age_1 uint8 = 31
	var age_2 = 32
	age_3 := 33
	fmt.Println(&age_1, &age_2)
	fmt.Println(age_1, age_2, age_3)
	sizeof := unsafe.Sizeof(age_3)
	fmt.Println(sizeof)

	var age_4, age_5, age_6 int = 31, 32, 33
	fmt.Println(age_4, age_5, age_6)

	var name_1, age_7 = "Tom", 30
	fmt.Println(name_1, age_7)

	name_2, is_boy, height := "Jay", true, 180.66
	fmt.Println(name_2, is_boy, height)
}
