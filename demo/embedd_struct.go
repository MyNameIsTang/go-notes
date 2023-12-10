package demo

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func InitEmbeddStruct() {
	b := B{A{2, 3}, 2.1, 4}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
