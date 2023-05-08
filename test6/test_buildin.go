package test6

import "fmt"

func TestBuildin() {
	// 1. 复数
	c1 := complex(1, 2)
	fmt.Println("c1=", c1)
	r1, i1 := real(c1), imag(c1)
	fmt.Println("r1=", r1, "i1=", i1)
}