package demo

import "fmt"

type Request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

func (r *Request) String() string {
	return fmt.Sprintf("a: %d, b: %d, a+b=%d", r.a, r.b, <-r.replyc)
}

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (chan *Request, chan bool) {
	reqChan := make(chan *Request)
	quit := make(chan bool)
	go server(op, reqChan, quit)
	return reqChan, quit
}

func InitMultipleServer() {
	adder, quit := startServer(func(a, b int) int { return a + b })

	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}

	adder <- req1
	adder <- req2

	fmt.Println(req1, "\n", req2)

	// const N = 100
	// var reqs [N]Request
	// for i := 0; i < N; i++ {
	// 	req := &reqs[i]
	// 	req.a = i
	// 	req.b = i + N
	// 	req.replyc = make(chan int)
	// 	adder <- req
	// }
	// for i := N - 1; i >= 0; i-- {
	// 	if <-reqs[i].replyc != N+2*i {
	// 		fmt.Println("fail at", i)
	// 	} else {
	// 		fmt.Println("Request ", i, "is ok!")
	// 	}
	// }
	quit <- true
	fmt.Println("done")
}
