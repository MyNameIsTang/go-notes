package demo

import (
	"fmt"
	"time"
)

func InitChannelBlock3() {
	ch1 := make(chan int, 10)

	go func() {
		time.Sleep(15 * 1e9)
		x := <-ch1
		fmt.Println("Sleep 15", x)
	}()

	fmt.Println("sending", 10)
	ch1 <- 10
	fmt.Println("send", 10)
}
