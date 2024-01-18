package demo

import "fmt"

type empty interface{}

var sum2 float64

type semaphore chan empty

func (s semaphore) sum1(a float64, b float64) {
	sum2 = a + b
}

func (s semaphore) lock() {
	e := new(empty)
	s <- e
}
func (s semaphore) unLock() {
	<-s
}

func InitGosum() {
	ch1 := make(semaphore, 1)
	ch1.lock()
	ch1.sum1(1, 2)
	fmt.Println("sum2:", sum2)
	ch1.unLock()
}
