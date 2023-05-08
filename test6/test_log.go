package test6

import (
	"GoNote/util"
	"errors"
)

func Testlog() {
	err := errors.New("可爱的错误")
	util.Error.Println(err)
	// util.Warning.Panic(err)
	util.Info.Fatal(err)
}