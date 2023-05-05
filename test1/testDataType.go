package test1

import(
	"fmt"
)

// 数据类型
func Testdatatype1(){
	var name string = "YQS"
	age := 18
	isMan := true
	fmt.Printf("name type: %T\n", name)
	fmt.Printf("age type: %T\n", age)
	fmt.Printf("isMan type: %T\n", isMan)
}

// 指针类型
func Testdatatype2(){
	num := 10
	// 指针类型
	var p *int = &num
	fmt.Printf("p type: %T\n", p)
	fmt.Printf("p value: %v\n", p)
	fmt.Printf("num address: %v\n", &num)
	fmt.Printf("p point value: %v\n", *p)
}

// 数组: 固定长度
func Testdatatype3(){
	array := [3]int{1, 2, 3}
	fmt.Printf("array type: %T\n", array)
}

// 切片: 动态数组
func Testdatatype4(){
	slice := []int{1, 2, 3, 4}
	fmt.Printf("slice type: %T\n", slice)
}
