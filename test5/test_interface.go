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
