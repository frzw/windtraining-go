package base

import "fmt"

func Defer() {
	defer fmt.Println("x")
	panic("1")
	fmt.Println("y")
	return
}
