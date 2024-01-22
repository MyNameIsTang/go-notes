package demo

import (
	"fmt"
)

// func fibonacci(n int) int {
// 	if n <= 1 {
// 		return 1
// 	}
// 	return fibonacci(n-1) + fibonacci(n-2)
// }

// func getData4(n int, ch chan int) {
// 	for i := 0; i < n; i++ {
// 		ch <- fibonacci(i)
// 	}
// 	close(ch)
// }

// func fibonacci2(n int, ch chan int) {
// 	x, y := 1, 1
// 	for i := 0; i < n; i++ {
// 		ch <- x
// 		x, y = y, x+y
// 	}
// 	close(ch)
// }

func fibonacci3(ch chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func InitGoFibonacci() {
	ch := make(chan int, 10)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("value is: ", <-ch)
		}
		done <- true
	}()
	fibonacci3(ch, done)

	// go fibonacci2(cap(ch), ch)
	// for {
	// 	select {
	// 	case v, ok := <-ch:
	// 		if ok {
	// 			fmt.Println("value is: ", v)
	// 		} else {
	// 			os.Exit(0)
	// 		}
	// 	default:
	// 		fmt.Println("waiting...")
	// 	}
	// }
	// for v := range ch {
	// 	fmt.Println("value is: ", v)
	// }

	// start := time.Now()
	// go getData4(10, ch)
	// for {
	// 	if v, ok := <-ch; ok {
	// 		fmt.Println("value is: ", v)
	// 	} else {
	// 		end := time.Now()
	// 		delta := end.Sub(start)
	// 		fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	// 		os.Exit(0)
	// 	}
	// }
}
