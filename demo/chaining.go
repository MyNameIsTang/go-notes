package demo

import (
	"flag"
	"fmt"
)

// 链式协程
var ngoroutine = flag.Int("n1", 100000, "how many goroutines")

func f2(left, right chan int) {
	v := <-right
	fmt.Println("i", v)
	left <- 1 + v
}

func InitChaining() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f2(left, right)
	}
	fmt.Println("cc")
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}
