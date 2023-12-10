package demo

import "fmt"

type amonymous1 struct {
	a float32
	int
	string
}

func InitAmonymousStruct() {
	a := amonymous1{2.3, 3, "哈哈哈"}
	fmt.Println(a)
}
