package base

import "fmt"

type Person struct {
	Name      string
	Age       int
	Sex       int
	HairCount int
	Salary    int
}

type Employee struct {
	EID   string
	EName string
	Person
}

func FangFa() {
	var erdanli Employee
	erdanli.person("lizhi", 23, 0, 12000, 100000)
	fmt.Println("person:", erdanli)
	erdanli.employee("10086", "erdanli")
	fmt.Println("employee:", erdanli)
	erdanli.addSalary("10086", 6000, 1000)
	fmt.Println(erdanli)
}

func (e *Employee) employee(eID, eName string) {
	e.EID = eID
	e.EName = eName
}

func (e *Employee) addSalary(eID string, salary, hairCount int) {
	e.EID = eID
	e.Salary += salary
	e.HairCount -= hairCount
}

func (p *Person) person(name string, age, sex, salary, hairCount int) {
	p.Name = name
	p.Age = age
	p.Sex = sex
	p.Salary = salary
	p.HairCount = hairCount
}

func (p Person) getAge() interface{} {
	return p.Age
}

func (p Person) getSex() interface{} {
	return p.Sex
}

func (p *Person) upAge() {
	p.Age++
}
