package demo

import "fmt"

type RSimple struct {
	v int
}

func (rs *RSimple) Get() int {
	return rs.v
}

func (rs *RSimple) Set(v int) {
	rs.v = v
}

func fi(si Simpler) {
	switch t := si.(type) {
	case *RSimple:
		fmt.Printf("Type RSimple %T with value %v\n", t, t)
	case *Simple:
		fmt.Printf("Type Simple %T with value %v\n", t, t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func gI(any Any) {
	switch t := any.(type) {
	case Simpler:
		fmt.Printf("Type Simpler %T with value %v\n", t, t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func InitSimpleInterface2() {
	var simple Simpler
	fi(simple)
	simple = &RSimple{4}
	simple.Set(23)
	fi(simple)
	simple = &Simple{34}
	fmt.Printf("Simple value is: %v\n", simple.Get())
	fi(simple)
	gI(simple)
	v := 12
	gI(v)
}
