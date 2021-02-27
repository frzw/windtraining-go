package base

import "fmt"

type ObjPhone struct {
	phone []Phone
}

type Phone interface {
	call()
}

type IPhone struct {
}

type Huawei struct {
}

func (i *IPhone) call() {
	fmt.Println("IPhone call, please !")
}

func (h *Huawei) call() {
	fmt.Println("Huawei call, please !")
}

func InterFace() {
	var objPhone ObjPhone
	var iPhone IPhone
	var huawei Huawei
	objPhone = ObjPhone{
		phone: []Phone{
			&iPhone,
			&huawei,
		},
	}
	for _, p := range objPhone.phone {
		p.call()
	}
	var phone Phone
	phone = new(IPhone)
	phone.call()

	phone = new(Huawei)
	phone.call()
}
