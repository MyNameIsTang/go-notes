package demo

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func InitMyTime() {
	m := myTime{time.Now()}
	fmt.Println(m.first3Chars())
}
