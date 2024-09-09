package main

import "fmt"

/**
包

在给一个文件打包时，该包对应一个文件夹， 比如这里的utils文件夹对应的包名就是utils。
文件的包名通常和文件所在的文件夹名一致，一般为小写


导包时
相对导入， 从当前目录下寻找

绝对导入， 从$GOPATH的src下开始，从其目录下寻找

包中的函数需要首字母大写，类似于public。这样才可以使用

包名较长时，可以给其取别名， 取完别名后，之前的包名就不能用了

同一个包下，不能有相同的函数名。也不能有相同的全局变量名


*/

func main() {
	var a int = int('1')
	b := int('2')
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T", a)
}
