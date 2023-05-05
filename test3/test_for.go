package test3

import (
	"fmt"
)

func TestFor() {
	fmt.Printf("for 无限循环\n")
	i := 0
	for {
		fmt.Printf("i = %v\n", i)
		i++
		if i > 10 {
			break
		}
	}
	fmt.Printf("for 条件循环\n")
	for i <= 20{
		fmt.Printf("i = %v\n", i)
		i++

	}

	fmt.Printf("for 标准循环\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %v\n", i)
	}

	
}