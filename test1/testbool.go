package test1

import(
	"fmt"
)

// bool 类型
func Testbool1(){
	var b1 bool = true
	var b2 bool = false
	fmt.Printf("b1 type: %T\n", b1)
	fmt.Printf("b2 type: %T\n", b2)
	b3 := true
	b4 := false
	fmt.Printf("b3 value: %v\n", b3)
	fmt.Printf("b4 value: %v\n", b4)
}

func Testbool2(){
	age := 16
	ok := age >= 18
	if ok {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}
}

func Testbool3(){
	//for 循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %v\n", i)
	}
}

func Testbool4(){
	age := 18
	gender := "男"
	if age >=18 && gender == "男" {
		fmt.Println("成年男性")
	} else {
		fmt.Println("未成年")
	}
}