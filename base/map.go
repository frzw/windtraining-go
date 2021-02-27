package base

import "fmt"

func Map() {
	// key: key 必须是支持==的数据类型
	// define 1.0
	var m map[int]int
	m = map[int]int{
		1: 1,
	}
	fmt.Println(m)

	// define 2.0
	m2 := make(map[int]int)
	m2[0] = 1
	m2[1] = 2
	fmt.Println(m2)
	for k, v := range m2 {
		fmt.Println(k, "->", v)
		delete(m2, k)
	}
	fmt.Println(m2)
}
