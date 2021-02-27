package base

import "fmt"

/* 函数
1.返回参数列表可以省略
2.可变参数：...type
*/
func HanShu() {
	sum1 := hanShu1_0(1, 2)
	fmt.Println(sum1)
	hanShu2_0(2, 3)
}

// 1.0
func hanShu1_0(x, y int) int {
	return x + y
}

// 2.0
func hanShu2_0(x, y int) {
	fmt.Println(x + y)
	return
}
