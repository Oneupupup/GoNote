package test2

import (
	"fmt"
)

func increment(n *int) {
	(*n)++ // n += 1
	fmt.Printf("结束时 n =  %v\n", *n)
	// n 的地址
	fmt.Printf("n 的地址 =  %v\n", n)
}

func Pionter() {
	var num int = 2023
	increment(&num)
	fmt.Printf("调用完成之后 num = %v\n", num)
	fmt.Printf("num 的地址 =  %v\n", &num)
	var ptr *int = new(int)
	fmt.Printf("ptr: %v\n", ptr)
	fmt.Printf("ptr point value: %v\n", *ptr)
	fmt.Printf("ptr address: %v\n", &ptr)
}
