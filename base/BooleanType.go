package base

import "fmt"

func BuErType() {
	// false true
	// 操作：&&（AND），||（OR）
	// 有效利用&&:可以优化代码性能
	b1 := false
	b2 := true
	if b1 && b2 {
		fmt.Println(true)
	}
	if b1 || b2 {
		fmt.Println(true)
	}
}
