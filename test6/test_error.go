package test6

import (
	"errors"
	"fmt"
)

func Errors1() {
	// 匿名函数 defer + recover
	defer func() {
		err := recover()
		fmt.Println("扑捉到了1个错误",err)
	}()
	// 自定义错误
	err1 := errors.New("err1")
	fmt.Println(err1)
	// 内置错误
	err2 := fmt.Errorf("%serr2","错误2")
	fmt.Println(err2)
	//发生恐慌
	panic(err2)
}