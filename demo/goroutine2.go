package demo

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}

func InitGoroutine2() {
	ch := make(chan string)
	go sendData(ch)
	time.Sleep(2e9)
	go getData(ch)
}
