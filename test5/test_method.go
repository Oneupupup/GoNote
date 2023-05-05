package test5

import (
	"fmt"
)

// User 结构体
type User struct {
	name string
	Id   uint32
}

// 小写的方法，只能在同一个包中使用
func (u User)printName(){
	fmt.Println("name = ", u.name)
}

func (u *User)setId(id uint32){
	u.Id = id
}

func TestMethod1() {
	u1 := User{
		name: "xiaofang",
		Id:   1000,
	}
	u1.setId(2000)
	u1.printName()
	fmt.Printf("u1: %v\n", u1)
}