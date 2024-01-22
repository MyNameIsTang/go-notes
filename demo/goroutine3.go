package demo

import "fmt"

func sendData3(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData3(ch chan string) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func InitGoroutine3() {
	ch := make(chan string)
	go sendData3(ch)
	getData3(ch)
}
