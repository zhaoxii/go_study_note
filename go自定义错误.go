package main

import (
	"errors"
	"fmt"
)

/*
*

go 支持自定义错误，使用errors.New 和panic内置函数。
errors.New("错误说明")，会返回一个error类型的值，表示一个错误。
panic内置函数，接受一个interface{}类型的值（也就是任何值了）作为参数。可以接受error类型的变量。输出错误信息并退出程序
*/

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("文件名称错误...")
	}
}

func test222() {
	err := readConf("confirg.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test222()")
}

func main() {
	test222()
	fmt.Println("main....")
}
