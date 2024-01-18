package demo

import (
	"fmt"
)

func send1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * 10
	}
	close(ch)
}

func get1(ch1 chan int, ch2 chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}
	ch2 <- true
}

func InitProducerConsumer() {
	ch := make(chan int)
	done := make(chan bool)
	go send1(ch)
	go get1(ch, done)
	<-done
}
