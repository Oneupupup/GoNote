package test4

import (
	"fmt"
)

// 自定义类型
func TestdefineType1() {
	// 自定义类型
	type MessageType uint16
	var u1000 uint16 = 1000
	var testMessage MessageType = MessageType(u1000)
	// print type
	fmt.Printf("testMessage: %T\n", testMessage)
	fmt.Printf("testMessage: %v\n", testMessage)
}

func TestdefineType2() {
	type myUnit16 = uint16
	var myu16 myUnit16 = 1000
	fmt.Printf("myu16: %T\n", myu16)
	fmt.Printf("myu16: %v\n", myu16)
}