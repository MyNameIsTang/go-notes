package demo

import "fmt"

type Address struct {
	number   int
	country  string
	city     string
	province string
}

type Birth struct {
	year  int
	month int
	day   int
}

type VCard struct {
	name    string
	birth   *Birth
	addr    *Address
	picture string
}

func InitVCard() {
	addr := &Address{number: 1, country: "中国", city: "北京", province: "顺义"}
	birth := &Birth{year: 1996, month: 2, day: 21}
	vCard := new(VCard)
	vCard.name = "汤姆"
	vCard.birth = birth
	vCard.addr = addr
	vCard.picture = "帅哥"
	fmt.Printf("%v", vCard)
}
