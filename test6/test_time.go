package test6

import (
	"fmt"
	"time"
)

func TestTime1() {
	// time 包
	// 时段
	// type Duration int64
	// sleep
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
		time.Sleep(time.Millisecond)
	}
	// 解析时间
	// ParseDuration(s string) (Duration, error)
	d1 ,err := time.ParseDuration("1h30m")
	if err != nil {
		panic(err)
	}else{
		fmt.Println("d1 = ", d1)
	}

	// Parse
	// Parse(layout, value string) (Time, error)
	d2, err := time.Parse("2006-01-02 15:04:05", "2023-08-08 08:08:08")
	if err != nil {
		panic(err)
	}else{
		fmt.Println(time.Since(d2))
	}
}

func TestTime2() {
	var intChan chan int = make(chan int)
	select{
		case<-intChan:
			fmt.Println("收到了验证码")
		case <- time.After(time.Second):
			fmt.Println("超时了")
	}
}

// 时区
func TestTime3() {
	l1,err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println(l1)
}

func TestTime4() {
	var intChan chan int = make(chan int)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	TickerFor:
	select{
		case<-intChan:
			fmt.Println()
			break TickerFor
		case <- time.NewTicker(time.Millisecond * 100).C:
			fmt.Println(".")
		}
}