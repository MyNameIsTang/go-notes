package demo

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func Perimeter(r *Rectangle) int {
	return (r.length + r.width) * 2
}

func InitRectangle() {
	rect := &Rectangle{10, 4}
	fmt.Printf("area is %v\n", Area(rect))
	fmt.Printf("perimeter is %v\n", Perimeter(rect))
}
