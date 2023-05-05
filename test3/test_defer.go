package test3

import (
	"fmt"
)

func deferUtil() func(int) int {
	i:=0
	return func(n int) int {
		fmt.Printf("n = %v\n", n)
		i++
		fmt.Printf("i = %v\n", i)
		return i
	}
}

func Defer() int {
	f := deferUtil()
	defer f(1)
	defer f(2)
	defer f(3)
	return f(4)
}

func DeferRecover() {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("err = %v\n", err)
		}
	}()
	n := 0
	fmt.Printf("10 / n = %v\n", 10 / n)
}