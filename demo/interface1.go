package demo

import "fmt"

type Shaper interface {
	Area() float32
	// Perimeter() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func InitInterface1() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper = sq1
	fmt.Println(areaIntf.Area())
	fmt.Printf("%#v", areaIntf)
}
