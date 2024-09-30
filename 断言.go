package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	fmt.Println("hello world")

	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	// b = a不可以
	b = a.(Point) // 接口断言进行类型转换。表示判断a是否是指向Point类型的变量，如果是就转成Point类型
	fmt.Println(b)

	var x interface{}
	var b2 float32 = 1.1
	x = b2
	y, ok := x.(float32)
	if ok == true {
		fmt.Println("success")
		fmt.Println(y)
	} else {
		fmt.Println("fail")
		fmt.Println(y)
	}

	fmt.Println("continue")

	TypeJudge(1, 1.1, x, y)

}

// TypeJudge 判断传入参数类型
func TypeJudge(items ...interface{}) {
	fmt.Println("typeJudge....................")
	for i, x := range items {
		fmt.Println("x=", x)
		switch x.(type) { //固定写法
		case Point:
			fmt.Println("Point:", i)
		case int, int8, int16, int32, int64:
			fmt.Println("int:", i)
		case float32, float64:
			fmt.Println("float32:", i)
		case string:
			fmt.Println("string:", i)
		case bool:
			fmt.Println("bool:", i)
		case nil:
			fmt.Println("nil:", i)
		default:
			fmt.Println("not found")
		}
	}
}
