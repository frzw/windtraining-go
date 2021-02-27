package base

import "fmt"

func WriteName() {
	// 下划线或字母开头
	var _lizhi string
	var lizhi_ int

	// 驼峰命名,不要蛇形命名 _li_zhi_wan_
	var lizhiData int

	// 字母大小写是不同得
	var lizhi int
	var Lizhi string

	// 命名要简短有意义
	var lifeIsAFuckingMovie string

	fmt.Print(lizhiData, lizhi, Lizhi, lifeIsAFuckingMovie, _lizhi, lizhi_)
}
