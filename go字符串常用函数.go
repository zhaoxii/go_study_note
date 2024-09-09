package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
1、len() 统计字符串的长度。按字节统计，不是字符。 go的编码统一为utf8（ascii的字符占1个字节，汉字站3个）
2、字符串遍历，同时处理有中文的问题。 r:=[]rune(str)
3、字符串转整数  strconv.Atoi("21")
4、整数转字符串  strconv.Itoa(1234)
5、字符串转 []byte
6、byte转字符串
7、10进制转其他进制strconv.FormatInt(123,2).。返回对应的字符串
8、查找子串 是否在指定的字符串中。 strings.Contains("xxxxxx", "xx")
9、统计一个字符串中包含几个指定的子串  strings.Count("xxxx", "xx")
10、不区分大小写的比较字符串是否相等。 strings.EqualFold("abc", "ABC")。 ==是区分大小写的
11、返回子串在字符串中第一次出现的位置，如果没有返回-1. strings.Index("asdaadsa", "s")


*/

func main() {
	fmt.Println("hello world")
	str := "zzzzz"
	fmt.Println(len(str))
	str1 := "赵熙xx" // go的编码统一为utf8（ascii的字符占1个字节，汉字站3个）
	fmt.Println(len(str1))

	str2 := "hello北京"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符=%c\n", str2[i]) // 按字节遍历的
	}
	str3 := []rune("hello北京")
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符=%c\n", str3[i]) // 按字符遍历
	}
	res, err := strconv.Atoi("121121")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", res)
	fmt.Printf("%T\n", strconv.Itoa(1111111111))

	var bytes = []byte("aaa")
	fmt.Println(bytes) // ascii 编码

	str4 := string([]byte{97, 98, 99})
	fmt.Printf("%v, %T\n", str4, str4)

	fmt.Println(strconv.FormatInt(123, 2))
	fmt.Println(strconv.FormatInt(123, 8))
	fmt.Println(strconv.FormatInt(123, 16))

	fmt.Println(strings.Contains("zxxzxzxczxcz", "zxc"))
	fmt.Println(strings.Count("xxxx", "xx"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))
	fmt.Println(strings.EqualFold("abc", "ABC"))
	fmt.Println(strings.Index("asdasda", "d"))
	fmt.Println(strings.Index("d阿达赵熙大苏打", "赵"))

	fmt.Println(strings.LastIndex("adsdaadas", "a"))
	fmt.Println(strings.LastIndex("d阿大达赵熙大苏打", "大"))

	fmt.Println(strings.Replace("zhaoxixxx", "x", "s", 1))
	fmt.Println(strings.Replace("zhaoxixxx", "x", "s", 2))
	fmt.Println(strings.Replace("zhaoxixxx", "x", "s", -1))

	for i, i2 := range strings.Split("xx,xxxx", ",") {
		fmt.Println(i, i2)
	}

	fmt.Println(strings.ToLower("ASDDDDDDD"), strings.ToUpper("xxxxxxxxxx"))
	fmt.Println(strings.TrimSpace(" xxxx xxx "))
	fmt.Println(strings.Trim("! DS!ADAS! !", "! "))
	fmt.Println(strings.TrimLeft("! DS!ADAS! !", "! "))
	fmt.Println(strings.TrimRight("! DS!ADAS! !", "! "))
	fmt.Println(strings.HasSuffix("@xczxcx", "@"))
	fmt.Println(strings.HasPrefix("@xczxcx", "@"))

}
