package demo

import "fmt"

type Day int

var weeks []string = []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}

const (
	_ = iota
	MO
	TU
)

func (d Day) String() string {
	return weeks[d-1]
}

func InitDay() {
	var d Day = 2
	fmt.Println(d)
}
