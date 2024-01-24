package demo

import (
	"fmt"
	"math"
	"time"
)

func CalcularePi(n int) float64 {
	ch1 := make(chan float64)
	for i := 0; i < n; i++ {
		go term(ch1, float64(i))
	}
	f := 0.0
	for i := 0; i < n; i++ {
		f += <-ch1
	}
	return f
}

func term(ch chan float64, i float64) {
	ch <- 4 * math.Pow(-1, i) / (2*i + 1)
}

func InitConcurrentPi() {
	start := time.Now()
	println(CalcularePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
