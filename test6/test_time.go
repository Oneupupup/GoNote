package test6

import(
	"fmt"
	"time"
)

func TestTime() {
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