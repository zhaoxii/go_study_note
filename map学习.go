package main

import "fmt"

/*
map是key-value的数据结构
基本语法：
var 名称 map[key类型]value类型。
key 可以是bool，数字，string，指针，chan，还可以是接口，结构体，数组。 slice，map，function不可以，因为没法用==判断
通常是int， string

value 的类型跟key一样。通常是string， 数字， map， struct

定义例子：
1、var a map[string]string
2、var b map[string]map[string]string

注意：map声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用


map 是引用类型, 遵循引用类型传递的方式.



*/

func main() {
	fmt.Println("hello world")
	var a map[string]string
	//a["zz"] = "asd"  // 需要先make分配内存
	a = make(map[string]string, 10)
	a["zz"] = "asd"
	a["z"] = "asd"
	a["zasd"] = "asd"
	a["zz"] = "asdxxx"
	fmt.Println(a)

	// map使用方式1
	var b map[string]string
	b = make(map[string]string, 10)
	fmt.Println(b)

	// map 使用方式2. 直接make
	var c = make(map[string]string)
	c["zz"] = "asd"
	fmt.Println(c)

	// 使用方式3
	var d = map[string]string{"a": "a"}
	d["c"] = "c"
	fmt.Println(d)

	// map 的增删改查操作
	fmt.Println("=====================================")
	var v = make(map[string]string)
	v["a"] = "a"
	v["c"] = "c"
	fmt.Println(v)
	delete(v, "a") // 使用内置函数delete，如果key存在就删除，不存在也不会报错
	fmt.Println(v)
	fmt.Println(v["c"])
	val, findRes := v["c"]
	fmt.Println(findRes, val) // findRes 如果有key则为true，否则为false
	// 清空map中所有的key时，没有一个专门的方法一次删除，可以遍历一下key，逐个删除。
	// 或者map =make(...), make一个新的，让原来的成为垃圾，让gc回收

	// map遍历
	var vvv = make(map[string]map[string]string)
	vvv["a"] = map[string]string{"安安": "先写下"}
	vvv["b"] = map[string]string{"b": "前五大区的"}
	fmt.Println(vvv, len(vvv))

	for k, va := range vvv {
		fmt.Println(k)
		for kk, vv := range va {
			fmt.Println(kk, vv)
		}
	}

	// map 切片。切片的数据类型是map，则我们称为map切片.map的个数就可以动态变化了
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "xxx"
		monsters[0]["age"] = "12"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "aaa"
		monsters[1]["age"] = "12"
	}

	fmt.Println(monsters)

	newMonster := map[string]string{"name": "sss", "age": "11"}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

	// 练习
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	userInfo := make(map[string]map[string]string)

	newUserInfo := map[string]string{"name": "asd", "pwd": "111"}
	userInfo["zx"] = newUserInfo
	userInfo["zzz"] = map[string]string{"name": "sss", "pwd": "1232erf"}
	_, www := userInfo["zzz"]
	if www {
		userInfo["zzz"]["pwd"] = "99999"
	} else {
		userInfo["zzz"] = map[string]string{"name": "sss", "pwd": "1232erf"}
	}
	fmt.Println(userInfo)

}
