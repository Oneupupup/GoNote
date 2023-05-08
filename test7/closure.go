package test7

import "fmt"

// 闭包
// 闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在
// 即使已经离开了自由变量的环境也不会被释放或者删除，
// 闭包是由函数和与其相关的引用环境组合而成的实体。

func closureFunc() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("i = %d\n", n)
		i++
		fmt.Printf("i = %d\n", i)
		return i
	}
}

func Closure() {
	f := closureFunc()
	f(2)
	f(4)
	f = closureFunc()
	f(6)
}
