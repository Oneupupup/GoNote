package test6

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 随机数
func TestStd1() {
	// 随机数种子
	seedNum := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		rand.NewSource(seedNum)
		fmt.Println("rand.Int() = ", rand.Intn(10))
		seedNum++
	}
}

// 字符转换

func TestStd2() {
	i1 := 65
	s1 := "gmail.com"
	s2 := fmt.Sprintf("%d@%s", i1, s1)
	fmt.Println("s2 = ", s2)
	var (
		i2 int
		s3 string
	)

	// 字符串的解析
	n, err := fmt.Sscanf(s2, "%d@%s", &i2, &s3)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("n = ", n)
	fmt.Println("i2 = ", i2)
	fmt.Println("s3 = ", s3)

	// int ——> string
	// 第二个参数：进制
	s4 := strconv.FormatInt(123, 4)
	fmt.Println("s4 = ", s4)
	u1, err := strconv.ParseUint(s4, 4, 0)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("u1 = ", u1)
}
