package test4

import (
	"fmt"
)

// 数组
func TestArray1() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Printf("a[0] = %v\n", a[0])
}

func TestArray2() {
	var a [3]int = [3]int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v] = %v\n", i, a[i])
	}
}

func TestArray3() {
	var a [3]int = [3]int{1, 2, 3}
	// i 是索引，v 是值
	for i, v := range a {
		fmt.Printf("a[%v] = %v\n", i, v)
	}
}

// 二维数组
// 二维数组的英文是：multidimensional array
func TestArray4() {
	var multiArray [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for i, v := range multiArray {
		// v 是一维数组
		for j, v2 := range v {
			fmt.Printf("multiArray[%v][%v] = %v\t", i, j, v2)
		}
		fmt.Printf("\n")
	}
}