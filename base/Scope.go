package base

import "log"

func ZuoYongYu() {
	// 语法块，花括号包裹的代码块
	// 语法快内的内容无法被外部使用

	for true {
		x := 1
		if x == 1 {
			log.Println(x)
			break
		}
	}
}
