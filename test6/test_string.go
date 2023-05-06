package test6

import (
	"fmt"
	"strings"
)

func TestString() {
	// 存在子串
	s1 := "hello world"
	s2 := "world"
	fmt.Printf("s1 contains s2 ? %v\n", strings.Contains(s1, s2))
	// 子串出现的位置
	fmt.Printf("s2 in s1 ? %d\n", strings.Index(s1, s2))
	// 替换
	// Replace(s, old, new string, n int) string
	s3 := strings.Replace(s1, s2, "golang", 1)
	fmt.Printf("s3 = %s\n", s3)
	// 重复
	s4 := strings.Repeat(s2, 3)
	fmt.Printf("s4 = %s\n", s4)
	// 前缀
	fmt.Printf("s1 has prefix 'hello' ? %v\n", strings.HasPrefix(s1, "hello"))
	// 后缀
	fmt.Printf("s1 has suffix 'world' ? %v\n", strings.HasSuffix(s1, "world"))
	// 拆分 Fields 返回一个 slice，包含所有的子串
	s5 := "he-l-l o-wo-rld"
	s6 := strings.Fields(s5)
	fmt.Printf("s6 = %v\n", s6)
	// 拆分 Split 返回一个 slice，包含所有的子串
	s7 := strings.Split(s5, "-")
	fmt.Printf("s7 = %v\n", s7)
	// SplitN(s, sep string, n int) []string
	s8 := strings.SplitN(s5, "-", 2)
	fmt.Printf("s8 = %v\n", s8)


	// 修剪
	// Trim(s, cutset string) string
	// TrimLeft(s, cutset string) string
	// TrimRight(s, cutset string) string
	// xiujian
	s9 := "  hello world  "
	s10 := strings.Trim(s9, " ")
	fmt.Printf("s10 = %s\n", s10)
}