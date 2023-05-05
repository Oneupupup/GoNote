package util

import "fmt"

func SelectByKey(test ...string) (key int) {
	for i, v := range test {
		fmt.Printf("i = %v, v = %v\n", i + 1, v)
	}
	fmt.Printf("请输入要执行的测试用例编号：")
	// go 中的输入语句
	fmt.Scanln(&key)
	return 
}