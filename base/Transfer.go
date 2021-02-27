package base

import "fmt"

/* 函数调用
有形参的需要提供实参
*/
func DiaoYong(x int) {
	sum := diGui(x)
	fmt.Println(sum)
}

/* 递归 */
func diGui(x int) int {
	if x <= 0 {
		return 0
	}
	return x + diGui(x-1)
}
