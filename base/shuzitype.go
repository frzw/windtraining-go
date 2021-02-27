package base

import "fmt"

func ZhengXing() {
	// 1. 整型：有符号、无符号
	// int8, int16, int32, int64, uint8, uint16, uint32, uint64, int ,uint (默认为int32, uint32)

	// 2. 浮点型
	// float32, float64

	var x int = 1
	var y uint = 2

	// 类型转化
	fmt.Print(x, "-")
	fmt.Print(int(y) + x)
}
