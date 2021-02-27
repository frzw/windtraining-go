package base

import "fmt"

// 全局变量
var global int

func Quanju() {
	// 局部变量
	var cur string
	fmt.Print(cur)
	// 使用全局变量
	global = 1
	fmt.Print(global)
}
