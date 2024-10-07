package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("========")
	fmt.Println(s)
	fmt.Println("========")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, sex string, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Println("struct has fields number:", num)
	for i := 0; i < num; i++ {
		fmt.Println("field:", val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println("tag:", tagVal)
		}
	}
	numofMethod := val.NumMethod()
	fmt.Println("struct has methods number:", numofMethod)
	val.Method(1).Call(nil) // 获取到第二个方法。 方法是按函数名排序的
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res:", res[0].Int(), res)
}

func main() {
	a := Monster{
		Name:  "xx",
		Age:   20,
		Score: 303.3,
	}
	TestStruct(a)
}
