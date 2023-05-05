package test5

import (
	"fmt"
)

func pushNum(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	// 关闭 channel
	close(ch)
}

// channel 通道
func TestChannel1(){
	// 声明 channel
	var ch chan int = make(chan int)
	// 启动协程
	go pushNum(ch)
	// 从 channel 中读取数据
	for num := range ch {
		fmt.Printf("num = %d\n", num)
	}
}