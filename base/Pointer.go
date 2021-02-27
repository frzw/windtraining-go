package base

import "fmt"

func Zhizhen() {
	// 指针是变量在内存中的地址

	// 定义
	var p *string

	x := "lizhi"

	p = &x

	fmt.Println(p)

	y := 9
	p2 := &y
	*p2 = 10
	fmt.Println(p2)

	fmt.Println(y)

	// 取指针指向得值一定要判断是否为空
	var a *int
	if a != nil {
		fmt.Println("a", *a)
	}
}
