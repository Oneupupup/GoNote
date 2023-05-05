package test1

import "fmt"

// 常量
// 常量的定义与变量类似，只不过使用 const 关键字。
// 定义就要赋值，不赋值编译不过去

func Testconst1(){
	// 语法：const identifier [type] = value
	const constname string = "YQS"
	const constage int = 18
	const constisMan bool = true
	// 打印类型
	fmt.Printf("constname type: %T\n", constname)
	fmt.Printf("constage type: %T\n", constage)
	fmt.Printf("constisMan type: %T\n", constisMan)
}

func Testconst2(){
	const PI float64 = 3.1415926
	const zero = 0.0 // 无类型浮点常量

	const (
		width = 100
		height = 200
	)

	const a, b = 3, 4

	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("zero: %v\n", zero)
	fmt.Printf("width: %v\n", width)
	fmt.Printf("height: %v\n", height)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

func Testiota1()  {
	// iota : 特殊常量，可以认为是一个可以被编译器修改的常量。
	// iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，
	// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}

func Testiota2()  {
	const (
		a = iota
		_
		b = iota
		_
		c = iota
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}

// iota声明中间插队
func Testiota3()  {
	const (
		a = iota
		b = 100
		c = iota
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}