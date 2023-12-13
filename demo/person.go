package demo

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

type number struct {
	f float32
}

type nr number

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func upPerson2(p Person) (res Person) {
	res.firstName = strings.ToUpper(p.firstName)
	res.lastName = strings.ToUpper(p.lastName)
	return
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(name string) {
	p.firstName = name
}

func InitPerson() {
	person := new(Person)
	person.firstName = "Chris"
	person.lastName = "Woodward"
	(*person).lastName = "Woodward"
	person2 := upPerson2(*person)
	fmt.Println(person2)

	// a := number{1.0}
	// b := nr{2.3}
	// c := number(b)
	// fmt.Println(a, b, c)
}
