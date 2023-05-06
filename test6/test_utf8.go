package test6

import (
	"fmt"
	_"strings"
	"unicode/utf8"
)

func Testutf8() {
	var s1 string = "你好，世界"
	fmt.Println(utf8.RuneCountInString(s1))
	// ValidString(s string) bool
	// ValidRune(r rune) bool
	// RuneLen(r rune) int
	// RuneCount(b []byte) int
	// EncodeRune(p []byte, r rune) int

}