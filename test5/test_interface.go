package test5

import (
	"fmt"
)

// 接口：interface
// 接口是一种类型，一种抽象的类型
// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口
type textMassage struct {
	Text string
}

func (tm *textMassage) setText() {
	tm.Text = "hello"
}

type imageMassage struct {
	image string
}

func (img *imageMassage) setImage() {
	img.image = "清明上河图"
}

// 定义一个接口
type Mes interface {
	setType()
}

func (tm *textMassage) setType() {
	tm.Text = "文字消息"
}
func (img *imageMassage) setType() {
	img.image = "图片消息"
}

func sendMes(m Mes) {
	m.setType()
	switch mptr := m.(type) {
	case *textMassage:
		mptr.setText()
	case *imageMassage:
		mptr.setImage()
	}
	fmt.Println("m = ", m)
}

func TestInterface1() {
	tm := textMassage{}
	img := imageMassage{}
	sendMes(&tm)
	sendMes(&img)
}

func TestInterface2(){
	var n1 int = 1
	// 空接口
	n1interface := interface{}(n1)
	// 空接口转换为其他类型
	n2 := n1interface.(int)
	fmt.Printf("n2 = %v\n",n2)
	// 空接口转换为其他类型，如果转换失败，会报错
	n3,ok := n1interface.(int)
	if ok {
		fmt.Printf("n3 = %v\n",n3)
	}else{
		fmt.Printf("n1interface can't convert to int\n")
	}
	
}
