package test3

import (
	"fmt"
)

func getSum(n1, n2 int) int {
	return n1 + n2
}

func getSumandDiff1(n1, n2 int) (sum ,diff int) {
	sum = n1 + n2
	diff = n1 - n2
	return sum, diff
}

func getSumandDiff2(n1, n2 int) (sum ,diff int) {
	sum = n1 + n2
	diff = n1 - n2
	// return sum, diff
	return
}

func Function() {
	res := getSum(1, 2)
	fmt.Printf("res = %v\n", res)
	sum, diff := getSumandDiff1(10, 20)
	fmt.Printf("sum = %v, diff = %v\n", sum, diff)
	sum, diff = getSumandDiff2(10, 20)
	fmt.Printf("sum = %v, diff = %v\n", sum, diff)

	// 打印函数地址和类型
	fmt.Printf("getSum type: %T, address: %p\n", getSum, getSum)

	// 匿名函数
	// 匿名函数没有函数名，只有函数体
	var getSum2 = func(n1, n2 int) int {
		return n1 + n2
	}
	res2 := getSum2(1, 2)
	fmt.Printf("getSum2 type: %T, address: %p\n", getSum2, getSum2)
	fmt.Printf("res2 = %v\n", res2)
}