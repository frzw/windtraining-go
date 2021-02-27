package base

import "fmt"

func BianLian() {
	// 命名
	var num1 int
	var num2 int = 12
	//  简单命名
	num := 1
	// 批量命名
	var i, j, k int
	var i1, j1, k1 = 1, 2, "skr"
	i2, j2, k2 := 1, 2, 3
	fmt.Print(num, num1, num2, i, j, k, i1, j1, j2, k1, i2, j2, k2)

	// 通过函数返回值初始化
	rV := returnValue()
	fmt.Print(rV)

}

func returnValue() int {
	return 1
}
