package demo

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	v float32
	int
	innerS
}

func InitStructAnonymouse() {
	outer := new(outerS)
	outer.b = 6
	outer.v = 1.2
	outer.int = 2
	outer.in1 = 2

	fmt.Println(outer)
}
