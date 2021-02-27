package base

import "fmt"

/* recover只defer中才生效 	*/
func Recover(num int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(num)
}
