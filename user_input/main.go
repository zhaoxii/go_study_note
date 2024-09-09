package main

import "fmt"

func main() {
	//var name string
	//var age int
	//var salary float64
	//var isPass bool
	//fmt.Println("请输入姓名: ")
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄: ")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水: ")
	//fmt.Scanln(&salary)
	//fmt.Println("请输入是否通过考试: ")
	//fmt.Scanln(&isPass)
	//
	//fmt.Printf("name is  %v \n age is %v \n salary is %v \n ispass is %t", name, age, salary, isPass)
	input()
}

func input() {
	var name string
	var age int
	var salary float64
	var isPass bool

	fmt.Println("请输入姓名，年龄，薪水，是否通过考试： 用空格隔开： ")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("name is  %v \nage is %v \nsalary is %v \nispass is %t", name, age, salary, isPass)

}
