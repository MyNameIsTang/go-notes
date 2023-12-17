package demo

import "fmt"

type Base1 struct {
	id int
}

func (b *Base1) Id() int {
	return b.id
}

func (b *Base1) SetId(v int) {
	b.id = v
}

type Person2 struct {
	FirstName, LastName string
	Base1
}

type Employee2 struct {
	salary float64
	Person2
}

func InitInheritMethods() {
	em := &Employee2{10000, Person2{"Tom", "jack", Base1{3}}}
	fmt.Printf("Employee2 id is:%v\n", em.id)
	em.SetId(4)
	fmt.Printf("Employee2 id is:%v\n", em.Id())
}
