package base

import "fmt"

func QiePian() {
	// 切片：代表边长的序列，元素类型相同
	// define: make
	slice := make([]int, 10)
	slice[0] = 1

	slice2 := []int{1}

	// arr := []int{1, 2, 3}
	for _, v := range slice2 {
		v++
		fmt.Println("v", v)
	}
	fmt.Println(slice2)
}
