package main

import (
	"fmt"

	a_go "github.com/frzw/algorithm-go/base"
	base "github.com/frzw/windtraining-go/base"
)

func main() {
	base.InterFace()
	arr := []int{4, 32, 34, 2, 1, 3, 55, 88, 33, 12}
	res := a_go.BubbleSortDesc(arr)
	fmt.Println(res)
}
