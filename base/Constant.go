package base

import "fmt"

type Weekday int

func ChangLiang() {
	// 编译期、运行期：编译后运行
	const (
		// li := "lizi" // error
		lz = 123
	)

	fmt.Println(lz)

	// itoa：常量中用
	const (
		start = iota // 常量中 0 开始赋值
		Sud
		_
		_

		Tues
		Fri
	)
	fmt.Println(Sud, Tues, Fri)

	// type
	var day Weekday
	day = 1
	fmt.Print(day)
}
