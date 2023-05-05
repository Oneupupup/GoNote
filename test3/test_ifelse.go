package test3
import (
	"fmt"
)

func TestIfElse() {
	fmt.Println("请输入年龄：")
	var age int
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("你是未成年人")
	} else if age > 60 {
		fmt.Println("你是老年人")
	} else{
		fmt.Println("你是成年人")
	}
}