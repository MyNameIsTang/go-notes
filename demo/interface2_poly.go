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

type Triangle struct {
	bottom, height float32
}

func (tr *Triangle) Area2() float32 {
	return (tr.height * tr.bottom) / 2
}

type PeriInterface interface {
	Perimeter() float32
}

func (sq *Square2) Perimeter() float32 {
	return sq.side * 4
}
func InitInterface2Poly() {
	r := &Rectangle2{2, 4}
	q := &Square2{4}
	shapes := []Shaper2{r, q}
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area2())
	}
	fmt.Printf("Square2 Perimeter result is %f\n", q.Perimeter())
	var shaper Shaper2 = &Triangle{10, 4}
	fmt.Printf("triangle area is %f\n", shaper.Area2())
}
