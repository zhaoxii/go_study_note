package main

import "fmt"

/*
数组是可以存放多个同一类型数据的类型，数组也是一种数据类型，
在go中，数组是值类型，遵循值类型的操作


切片slice
切片是数组的一个引用，因此切片是引用类型，进行传递时，遵循引用类型的传递机制。
切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度都一样
切片的长度是可以变化的。因此切片是一个可以动态变化的数组
切片定义的基本语法：var 变量名 []类型。 比如： var a []int

*/

func main() {

	//	 定义一个数组， 默认值为零值(数值类型是0，字符串是"", bool是false)。 长度固定
	var hens = [6]float64{}
	hens[0] = 11
	fmt.Println(hens)

	// 数组的地址是第一个元素的地址
	fmt.Printf("%p\n", &hens)
	fmt.Println(&hens[0])
	fmt.Println(&hens[1]) // 第二个元素的地址是 第一个元素的地址基础上加上该类型占用的字节数。 所以数组的地址是连续的

	// 定义数组的几种方式：
	var array1 [3]int = [3]int{1, 2, 3}
	var array2 = [3]int{1, 2, 3}
	var array3 = [...]int{1, 2, 3}                          // 根据初始化时的元素个数判断数组大小
	var array4 = [3]string{1: "tom", 0: "jack", 2: "marry"} // 可以指定位置
	fmt.Println(array1, array2, array3, array4)

	// 数组的遍历方式for-range. index是下标，value是值。 index，value都是只在循环内可见的局部变量
	for index, value := range array1 {
		fmt.Println(index, value)
	}

	array5 := [26]byte{}
	for index, _ := range array5 {
		array5[index] = 'a' + byte(index)
		fmt.Printf("%c\n", array5[index])
	}
	fmt.Println(array5)

	// 切片的基本使用
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	//定义一个切片。引用到intArr数组的2-3个元素。 [1:3]表示 从下标1开始，到下标3结尾。但是不包含3
	slice1 := intArr[1:3]
	fmt.Println(intArr)
	fmt.Println(slice1, len(slice1), cap(slice1)) // len是长度(元素个数)，cap是容量。
	slice1[0] = 333
	fmt.Println(slice1, intArr)

	// 切片的三种定义方式
	//1、 slice1 := intArr[1:3]  定义一个切片，然后让切片去引用一个已经创建好的数组
	//2、 通过make创建切片 var 名称 []类型 = make([], len, [cap])。len是长度，cap是容量(可选参数)
	//3、定义一个切片，直接就指定具体数组，使用原理类似与make。 var slice2 []string =[]string{"a", "bb", "ccc"}
	var slice2 []int = make([]int, 4, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))
	slice2[0] = 100
	slice2[1] = 200
	fmt.Println(slice2)
	slice2 = append(slice2, 1000, 100000, 20) // append 内置函数，可以对切片进行动态追加
	slice2 = append(slice2, slice1...)        // append  追加切片. 三个点是必须的
	fmt.Println("qp", slice2)
	var slice3 []string = []string{"a", "bb", "a只写"}
	fmt.Println(slice3)

	// 切片遍历
	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}

	for index, value := range slice3 {
		fmt.Println(index, value)
	}

	// 切片的拷贝。 copy 两个参数类型要求都是切片
	var a []int = []int{1, 2, 3}
	var sliceCopy = make([]int, 10)
	fmt.Println(a)
	fmt.Println(sliceCopy)
	copy(sliceCopy, a)
	fmt.Println(a)
	fmt.Println(sliceCopy)

}
