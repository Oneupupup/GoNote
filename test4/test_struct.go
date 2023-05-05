package test4

import (
	"fmt"
)

func TestStruct1() {
	type User struct {
		name string
		Id   uint32
	}
	// struct 的实例化
	var u1 User
	u1.name = "张三"
	u1.Id = 1000
	fmt.Println("u1 = ", u1)
}

func TestStruct2() {
	type User struct {
		name string
		Id   uint32
	}

	// 结构体指针

	var u1 *User = &User{
		name: "张三",
		Id:   1000,
	}

	// go 语言中，可以直接使用结构体指针访问结构体的成员
	// (*u1).name = "李四"
	u1.name = "李四"
}

// 继承
func TestStruct3() {
	type User struct {
		name string
		Id   uint32
	}

	type Account struct {
		User
		Password string
	}

	type Contact struct {
		*User
		Remark string
	}

	// 创建一个Account的实例
	var a1 Account = Account{
		User: User{
			name: "张三",
			Id:   1000,
		},
		Password: "123456",
	}

	var c1 *Contact = &Contact{
		User: &User{
			name: "李四",
			Id:   1001,
		},
		Remark: "朋友",
	}
	fmt.Println("a1 = ", a1)
	fmt.Println("c1 = ", c1)
	fmt.Println("c1.User = ", c1.User)
}
