package demo

import (
	"fmt"
	"strconv"
)

type Person3 struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson3(name string, salary float64) *Person3 {
	p := &Person3{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person3) backend() {
	for f := range p.chF {
		f()
	}
}

func (p *Person3) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

func (p *Person3) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person3) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func InitConcAccess() {
	bs := NewPerson3("Smith Bill", 2503.23)
	fmt.Println(bs)
	bs.SetSalary(5403.23)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
