package main

import "fmt"

func main() {

	key := ""

	balance := 10000.0
	money := 0.0
	note := ""
	detail := "收支\t账户金额\t收支金额\t说明"

	for {
		fmt.Println("---------------家庭收支记账软件-------------")
		fmt.Println("---------------1、收支明细-------------")
		fmt.Println("---------------2、登记收入-------------")
		fmt.Println("---------------3、登记支出-------------")
		fmt.Println("---------------4、退出-------------")

		fmt.Println("请输入1-4：")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n-------------当前收支明细记录---------------")
			fmt.Println(detail)
		case "2":
			fmt.Println("本次收入金额")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			balance += money
			detail += ""
		case "3":
			fmt.Println("3")
		case "4":
			fmt.Println("退出...")
			return
		default:
			fmt.Println("请输入正确的选项！！！")
		}
	}

}
