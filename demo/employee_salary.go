package demo

import "fmt"

type employee struct {
	salary float32
}

func (e *employee) giveRaise(b float32) {
	e.salary += e.salary * b
}

func InitEmployee() {
	em := employee{1000}
	em.giveRaise(0.3)
	fmt.Println(em.salary)
}
