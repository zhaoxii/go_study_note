package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体（值类型）
一种自定义的数据类型，代表一类事物

声明结构体：
type 结构体名称 struct {
     field1 type
     field2 type
}
字段的类型可以是 基本类型，数组，引用类型
结构体的字段默认是零值。指针，slice，map的零值都是nil，即还没有分配内存。如果需要这样的字段，需要先make在使用

不同结构体的变量是独立，互不影响的。一个结构体变量字段改变，不会影响到别的。

结构体的所有字段在内存中是连续分布的

结构体是用户定义的类型，和其他类型进行转换时需要有完全相同的字段（名称，个数，类型）


结构体进行type重新定义（取别名），go认为是不同的数据类型，但是可以相互转换。
type Student struct{
	name string
}
type stu Student


结构体的每个字段上，可以写一个tag，该tag可以通过反射机制获取，常见的使用场景是序列化和反序列化

结构体类型是值类型，在方法调用中，遵循值类型的传递机制。值拷贝传递
如果希望在方法中修改结构体变量的值，可以通过结构体指针的方式处理



*/

type Cat struct {
	Name  string
	Age   int
	Color string
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	fmt.Println("hello world")
	var cat1 Cat
	cat1.Name = "xb"
	cat1.Age = 18
	cat1.Color = "white"
	fmt.Println(cat1)

	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("p1.ptr is nil")
	}
	if p1.slice == nil {
		fmt.Println("p1.slice is nil")
	}
	if p1.map1 == nil {
		fmt.Println("p1.map1 is nil")
	}

	p1.slice = make([]int, 3)
	p1.slice[0] = 100
	fmt.Println(p1)

	p1.map1 = make(map[string]string)
	p1.map1["zxz"] = "zx"
	fmt.Println(p1)

	a := 1
	p1.ptr = &a
	fmt.Println(p1)

	p2 := p1 // 值拷贝
	p2.Name = "zx"
	fmt.Println(p1)
	fmt.Println(p2)

	//创建结构体变量和访问结构体字段
	//1、var person Person
	//2、var person Person = Person{}
	//3、var person *Person = new(Person)
	//4、var person *Person =&Person{}
	// 3,4两种是 结构体指针。结构体指针访问字段的方式是 (*person).Name。但go做了简化，也可以person.Name.编译器会加上*
	var p3 Person = Person{Name: "zx", Age: 11, Scores: [5]float64{1.1, 2.21, 3.33331, 4, 5},
		ptr: new(int), slice: make([]int, 3), map1: map[string]string{"name": "zx"}}
	fmt.Println(p3)
	var person *Person = new(Person)
	(*person).Name = "zx" // 也可以 person.Name = "zx" 这样写。 go设计者为了使用方便。底层会对person.Name = "zx"处理，加上*，（*person）
	person.Name = "zxx"
	(*person).Age = 1121
	person.Age = 1222222
	fmt.Println(*person)

	fmt.Println("============================================")
	m1 := Monster{Name: "牛魔王", Age: 100, Skill: "吐血"}
	m1_str, err := json.Marshal(m1)
	fmt.Println(err)
	fmt.Println(string(m1_str))

}
