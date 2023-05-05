package test4

import (
	"fmt"
	"GoNote/util"
)

// slice 切片
// slice 的默认类型nil
func TestSlice1() {
	// 切片没有长度
	var arr []int = []int{1, 2, 3, 4, 5}
	// [start:end]，包含 start，不包含 end
	var subarr []int = arr[1:3]		// [2, 3]
	// [0:len(arr)] = [:]
	var subarr2 []int = subarr[1:2]	// [3]

	fmt.Printf("subarr2[0]: %v\n", subarr2[0])

	// 修改切片的类型
	var s[] int
	fmt.Println("does s == nil ?", s == nil)

	s = make([]int, 3, 5)	// len(s) = 3, cap(s) = 5
	// s = make([]int, 3)	// len(s) = 3, cap(s) = 3
	// 如果不写 cap，那么 cap = len
	fmt.Println("s = ", s)
}

func TestSlice2()  {
	var s[] int = make([]int,3,5)
	fmt.Println("s = ",s)
	fmt.Println("len(s) = ",len(s))
	fmt.Println("cap(s) = ",cap(s))

	// 由系统创建底层的数组
	s1 := []int{1,2,3,4,5}
	fmt.Println("s1 = ",s1)
	fmt.Println("len(s1) = ",len(s1))
	fmt.Println("cap(s1) = ",cap(s1))

	// 追加元素
	var s2[] int = append(s1,6,7,8)
	fmt.Println("s2 = ",s2)
	fmt.Printf("s1: %v\n", s1)
}

func TestSlice3()  {
	// 切片的拷贝
	var s1[] int = []int{1,2,3,4,5}
	var s2[] int = make([]int,5)
	copy(s2,s1)
	fmt.Println("s1 = ",s1)
	fmt.Println("s2 = ",s2)

	// 容量不够，只拷贝前3个元素
	var s3[] int = make([]int, 3)
	copy(s3, s1)
	fmt.Println("s3 = ",s3)
}

// string 和 slice 的区别
func TestSlice4()  {
	var str string = "hello 世界"
	fmt.Printf("str: %v\n", str)
	// 按字符串输出
	fmt.Printf("[]byte(str): %s\n", []byte(str))

	// 遍历字符串
	for i, v := range str {
		fmt.Printf("str[%v] = %c\n", i, v)
	}
}

func TestSlice5() {
	key := util.SelectByKey("注册", "登录", "退出")
	fmt.Printf("接收到 key = %v\n", key)
} 