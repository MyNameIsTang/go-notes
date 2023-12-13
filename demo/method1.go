package demo

import "fmt"

type TwoInts struct {
	a, b int
}

type IntVector []int

func (tn *TwoInts) Sum() int {
	return tn.a + tn.b
}

func (tn *TwoInts) Sum2(c int) int {
	return tn.a + tn.b + c
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func InitMethod1() {
	two1 := new(TwoInts)
	two1.a = 2
	two1.b = 10
	fmt.Println(two1.Sum())
	fmt.Println(two1.Sum2(2))

	int1 := IntVector{1, 3}
	fmt.Println(int1.Sum())
}
