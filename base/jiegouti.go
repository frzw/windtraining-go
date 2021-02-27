package base

import "fmt"

type Family struct {
	name   string
	age    int
	sex    int
	salary int
}

type HunanFamily struct {
	name   string
	family []Family
}

func JieGouTi() {
	// 结构体: struct、零个或多个任意类型值的集合、值称为结构体的成员
	// use 1.0
	var erDanFamily Family
	erDanFamily.name = "lizhi"
	erDanFamily.age = 23
	erDanFamily.sex = 0
	erDanFamily.salary = 12000
	fmt.Println(erDanFamily)

	// use 2.0
	eDFamily := &Family{
		name:   "liziwan",
		age:    22,
		sex:    1,
		salary: 5500,
	}
	fmt.Println(*eDFamily)

	pFamily := eDFamily
	pFamily.name = "lixuanji"
	fmt.Println(*eDFamily)
	fmt.Println(*pFamily)

	// HunanFamily
	hnPeople := HunanFamily{
		name:   "zhu zhou",
		family: []Family{*eDFamily, erDanFamily},
	}
	hnPeople.family[0].name = "lijianyu"
	hnPeople.family[1].name = "helanying"
	fmt.Println(hnPeople)
}
