package base

import "fmt"

func ShuZu() {
	// define
	// var arr [3]int

	// init 1.0
	// arr = [...]int{1, 2, 3}

	// init 2.0
	arr := [...]int{0: 2, 3: 8}
	v := arr[:]
	fmt.Println("v->", v)
	// visit 1.0
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[2])

	// visit 2.0
	for i := 0; i < len(arr); i++ {
		fmt.Println("v.2.0:", arr[i])
	}

	// visit 3.0
	for k, v := range arr {
		fmt.Println("v.3.0:", k, "->", v)
	}

	//

}
