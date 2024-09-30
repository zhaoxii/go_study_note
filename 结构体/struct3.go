package main

import "fmt"

/*
如果一个结构体里嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法。从而实现了多重继承

如果嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分

(建议不使用多重继承)

*/

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

func main() {
	fmt.Println("hello world")

	tv1 := TV{}
	fmt.Println(tv1)

	tv2 := TV{Goods{Name: "zxcz", Price: 100}, Brand{Name: "xzcz", Address: "dada"}}
	fmt.Println(tv2)
	tv2.Goods.Name = "dalei"
	fmt.Println(tv2)

}
