package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值。
func updateValue(b interface{}) {
	v := reflect.ValueOf(b)
	fmt.Println(v.Kind())
	fmt.Println(v.Elem().Kind())
	fmt.Printf("%T\n", v)
	fmt.Printf("%T\n", v.Elem())

	v.Elem().SetInt(33)
	fmt.Println(v)
	fmt.Println(b)

}

func main() {
	fmt.Println("hello world")
	num := 22
	fmt.Println(&num)
	updateValue(&num) // 必须传地址过去。然后调用Elem()方法

	fmt.Println(num)
}
