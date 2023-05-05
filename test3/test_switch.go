package test3

import (
	"fmt"
)

func TestSwitch() {
	fmt.Println("请输入年龄：")
	var age int
	fmt.Scanln(&age)
	switch  {
	case age < 18:
		fmt.Println("你是未成年人")
	case age > 60:
		fmt.Println("你是老年人")
	case age <=60 && age >= 18:
		fmt.Println("你是成年人")
	default:
		fmt.Println("输入错误")
	}
}