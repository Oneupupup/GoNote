package test2

import (
	"fmt"
)

func FmtVerbs1(){
	fmt.Println("Hello World")
	fmt.Printf("%%hello Golang\n")

	var n1 = 63
	// unicode码值
	fmt.Printf("n1 = %U\n", n1)
	// 二进制
	fmt.Printf("n1 = %b\n", n1)
	// 十进制
	fmt.Printf("n1 = %d\n", n1)
	// 八进制
	fmt.Printf("n1 = %o\n", n1)
	// 十六进制
	fmt.Printf("n1 = %x\n", n1)
	// 十六进制大写
	fmt.Printf("n1 = %X\n", n1)
}

//浮点数
func FmtVerbs2(){
	var f1 = 3.1415926
	fmt.Printf("f1 = %f\n", f1)
	fmt.Printf("f1 = %.2f\n", f1)
	fmt.Printf("f1 = %9.2f\n", f1)
	fmt.Printf("f1 = %9.f\n", f1)

	// 科学计数法
	fmt.Printf("f1 = %b\n", f1)
	fmt.Printf("f1 = %e\n", f1)
	fmt.Printf("f1 = %E\n", f1)
	fmt.Printf("f1 = %x\n", f1)
	fmt.Printf("%t\n", f1==3.1415926)

	// 字符串
	var str = "Hello World"
	fmt.Printf("str = %s\n", str)
	fmt.Printf("str = %q\n", str)
	fmt.Printf("str = %x\n", str)
	fmt.Printf("str = %X\n", str)
}

// 指针
func FmtVerbs3(){
	var num int = 10
	fmt.Printf("num address: %v\n", &num)
	fmt.Printf("num address: %p\n", &num)
}