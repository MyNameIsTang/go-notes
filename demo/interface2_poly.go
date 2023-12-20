package demo

import "fmt"

type Shaper2 interface {
	Area2() float32
}

type Square2 struct {
	side float32
}

func (sq *Square2) Area2() float32 {
	return sq.side * sq.side
}

type Rectangle2 struct {
	width, height float32
}

func (re *Rectangle2) Area2() float32 {
	return re.width * re.height
}

func InitInterface2Poly() {
	r := &Rectangle2{2, 4}
	q := &Square2{4}
	shapes := []Shaper2{r, q}
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area2())
	}
}
