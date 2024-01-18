package demo

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck1(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func InitChannelBlock() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck1(ch1)
	time.Sleep(1e9)
}
