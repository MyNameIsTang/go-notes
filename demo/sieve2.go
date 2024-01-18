package demo

import "fmt"

func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter2(ch chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-ch; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve2() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}

func InitSieve2() {
	ch := sieve2()
	for {
		fmt.Println(<-ch)
	}
}
