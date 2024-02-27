package demo

import "fmt"

const MAXREQS = 100

var sem = make(chan int, MAXREQS)

func process1(r *Request) {
	// dsa
	fmt.Println(r.a, r.b)
}

func handle1(r *Request) {
	sem <- 1
	fmt.Println("ding")
	process1(r)
	<-sem
}

func server1(service chan *Request) {
	for {
		request := <-service
		go handle1(request)
	}
}
func InitMaxtasks() {
	service := make(chan *Request)
	go server1(service)

	// var reqs [MAXREQS]Request
	// for i := 0; i < MAXREQS; i++ {
	// 	req := &reqs[i]
	// 	req.a = i
	// 	req.b = i + MAXREQS
	// 	req.replyc = make(chan int)
	// 	req.replyc <- i
	// 	service <- req
	// }

	// for i := MAXREQS - 1; i >= 0; i-- {
	// 	if <-reqs[i].replyc != MAXREQS {
	// 		fmt.Println("fail at", i)
	// 	} else {
	// 		fmt.Println("Request ", i, "is ok!")
	// 	}
	// }

}
