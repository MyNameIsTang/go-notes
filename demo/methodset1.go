package demo

import "fmt"

type List []int

func (l List) Len() int { return len(l) }
func (l *List) Append(val int) {
	*l = append(*l, val)
}

func InitMethodSet1() {
	var lst List
	lst.Append(1)
	fmt.Println(lst, lst.Len())

	plst := new(List)
	plst.Append(3)
	fmt.Println(plst, plst.Len())
}
