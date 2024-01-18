package demo

import (
	"fmt"
)

func f11(in chan int) {
	fmt.Println(<-in)
}

func InitBlocking() {
	out := make(chan int)
	out <- 2
	go f11(out)
}
