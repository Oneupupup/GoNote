package test4

import (
	"fmt"
)

func TestMap1() {
	var m0 map[string]int
	fmt.Println("m0 == nil ?", m0 == nil)

	var m1 map[string]string = make(map[string]string, 2)
	// 初始化默认 size = 1
	m1["早上"] = "起床"
	m1["中午"] = "吃饭"
	m1["晚上"] = "睡觉"
	fmt.Println("m1 = ", m1)

	m2 := map[string]string{
		"早上": "起床",
		"中午": "吃饭",
		"晚上": "睡觉",
		"夜里": "做梦",
	}

	fmt.Println("m2 = ", m2)
}

func TestMap2() {
	m2 := map[string]string{
		"早上": "起床",
		"中午": "吃饭",
		"晚上": "睡觉",
		"夜里": "做梦",
	}
	// map的查找
	// 返回两个值，第一个是key对应的value，第二个是key是否存在
	// 如果不存在，返回value对应类型的零值
	v,ok := m2["早上"]

	if ok {
		fmt.Println("m2[早上] = ", v)
	} else {
		fmt.Println("m2[早上] 不存在")
	}
}

// 删除map中的元素
func TestMap3() {
	m2 := map[string]string{
		"早上": "起床",
		"中午": "吃饭",
		"晚上": "睡觉",
		"夜里": "做梦",
	}
	// 删除map中的元素
	delete(m2, "夜里")
	fmt.Println("m2 = ", m2)
}

// map的遍历
func TestMap4() {
	m2 := map[string]string{
		"早上": "起床",
		"中午": "吃饭",
		"晚上": "睡觉",
		"夜里": "做梦",
	}
	for k,v := range m2 {
		fmt.Printf("m2[%s] = %s\n", k, v)
	}
}