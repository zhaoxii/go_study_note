package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	c := reflect.TypeOf(b)
	fmt.Println(c)
	fmt.Println("===============")
	fmt.Println(c.String())
	fmt.Println("===============")
	fmt.Printf("type:%T\n", c)

	d := reflect.ValueOf(b)
	fmt.Println(d)
	fmt.Printf("%T\n", d.Interface())
	fmt.Printf("%T\n", d)

	switch d.Interface().(type) { //固定写法
	case Student:
		fmt.Println("Student:", d.Interface().(Student))
	case int:
		fmt.Println("int:", d.Interface().(int))
	case float64:
		fmt.Println("float32:", d.Interface().(float64))
	case string:
		fmt.Println("string:", d.Interface().(string))
	case bool:
		fmt.Println("bool:", d.Interface().(bool))
	case nil:
		fmt.Println("nil:", d.Interface())
	default:
		fmt.Println("not found")
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {

	a := 31
	//res := reflect.ValueOf(a)
	//fmt.Println(res.Kind())
	//fmt.Println(res.Type())
	//b := reflect.TypeOf(res.Interface())
	//fmt.Println(b)
	reflectTest(a)

	stu := Student{Name: "ad", Age: 20}
	reflectTest(stu)
}
