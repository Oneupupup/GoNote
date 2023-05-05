package test1

import "fmt"

func TestVar1() {
	// 声明变量的一般形式是使用 var 关键字：
	// var identifier type
	// 声明变量一定要使用，不使用编译不过去
	//int
	var age int = 18
	// string
	var name string = "YQS"
	// bool
	var isMan bool = true
	// 打印类型
	fmt.Printf("age type: %T\n", age)
	fmt.Printf("name type: %T\n", name)
	fmt.Printf("isMan type: %T\n", isMan)


	// 批量申请
	var (
		age1 int = 18
		name1 string = "YQS"
		isMan1 bool = true
	)
	// 打印值
	fmt.Printf("age1: %v\n", age1)
	fmt.Printf("name1: %v\n", name1)
	fmt.Printf("isMan1: %v\n", isMan1)
}

func TestVar2(){
	// 类型推导
	var age = 18
	var name = "YQS"
	var isMan = true

	fmt.Printf("age type: %T\n", age)
	fmt.Printf("name type: %T\n", name)
	fmt.Printf("isMan type: %T\n", isMan)
}

func TestVar3(){
	// 批量初始化
	var name, age, isMan = "YQS", 18, true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("isMan: %v\n", isMan)
}

func TestVar4(){
	// 简短声明：只能在函数内部使用
	// 语法：name := value
	name, age, isMan := "YQS", 18, true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("isMan: %v\n", isMan)
}

func TestVar5()(string, int, bool){
	// 匿名
	// _ 匿名变量，丢弃数据不处理
	// 用于函数返回值，表示不处理函数返回值\
	return "YQS", 18, true
}