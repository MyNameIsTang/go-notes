package demo

import "fmt"

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func InitLazyEvalutaion() {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}
