package test6

import (
	"errors"
	"fmt"
)

func Errors1() {
	err1 := errors.New("err1")
	fmt.Println(err1)
	err2 := fmt.Errorf("%serr2","错误2")
	fmt.Println(err2)
	panic(err2)
}
