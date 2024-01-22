package demo

import "fmt"

func tel(ch chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- true
}

func InitGoroutineSelect() {
	ch := make(chan int)
	done := make(chan bool)
	go tel(ch, done)
	for {
		select {
		case v := <-ch:
			fmt.Println("v: ", v)
		case <-done:
			fmt.Println("Done")
			return
		default:
			fmt.Println("default")
		}
	}
}
