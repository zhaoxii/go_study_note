package main

import "fmt"

/**

函数内定义的变量叫局部变量，作用域仅限于函数内部。
函数外部定义的变量叫全局变量，作用域在整个包都有效，如果其首字母为大写，则作用域在整个程序都有效。
如果变量在一个代码块，比如for、if中，那么这个变量的作用域就在该代码块。

*/

// Name := 20 --> 错误的，该语句等价于  Var Name string; Name=20; 不能有赋值语句。赋值语句不能在函数体外
var Name = 20 // 这个不是赋值，属于初始化

func main() {
	fmt.Println("hello world")
	fmt.Println(Name)
}
