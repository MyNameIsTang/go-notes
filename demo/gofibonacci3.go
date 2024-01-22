package demo

import (
	"fmt"
	"time"
)

func sub(in <-chan int) (a, b, c chan int) {
	a, b, c = make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			v := <-in
			a <- v
			b <- v
			c <- v
		}
	}()
	return
}

func fib() <-chan int {
	ch := make(chan int, 2)
	a, b, out := sub(ch)
	go func() {
		ch <- 0
		ch <- 1
		<-a
		for {
			ch <- <-a + <-b
		}
	}()
	<-out
	return out
}

func InitFibonacci3() {
	start := time.Now()
	f := fib()
	for i := 0; i < 10; i++ {
		println(<-f)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("longCalculation took this amount of time: ", delta)
}
