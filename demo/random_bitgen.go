package demo

import (
	"fmt"
	"math/rand"
	"time"
)

func random1() <-chan int {

	out := make(chan int)

	fmt.Print("dasda")
	go func() {
		fmt.Print("12312312")
		rand.Seed(time.Now().UnixNano())
		for {
			if rand.Float64() < 0.5 {
				out <- 0
			} else {
				out <- 1
			}
		}
	}()
	return out
}

func InitRandomBitgen() {
	ch1 := random1()
	for {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Print(v)
			} else {
				return
			}
		}
	}
}
