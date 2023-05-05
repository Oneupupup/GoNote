package test4

import (
	"fmt"
)

func TestMap1(){
	var m0 map[string]int
	fmt.Println("m0 == nil ?", m0 == nil)

	var m map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("m[\"a\"] = %v\n", m["a"])
}