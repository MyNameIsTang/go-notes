package demo

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

func (c *Car) Start() {
	fmt.Println("car is starting")
}

func (c *Car) Stop() {
	fmt.Println("car is stoped")
}

func (m *Mercedes) sayHiToMerkel() {
	m.Start()
	fmt.Println("sayHiToMerkel")
	m.Stop()
}

func InitInheritanceCar() {
	m := new(Mercedes)
	m.wheelCount = 4
	fmt.Println(m.numberOfWheels())
	m.sayHiToMerkel()
}
