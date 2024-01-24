package demo

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const NCPU = 2

func CalcularePi2(n int) float64 {
	ch1 := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term2(ch1, float64((i*n)/NCPU), float64((i+1)*n/NCPU))
	}
	f := 0.0
	for i := 0; i < NCPU; i++ {
		f += <-ch1
	}
	return f
}

func term2(ch1 chan float64, start, end float64) {
	result := 0.0
	for i := start; i < end; i++ {
		result += 4 * math.Pow(-1, i) / (2*i + 1)
	}
	ch1 <- result
}

func InitConcurrentPi2() {
	start := time.Now()
	runtime.GOMAXPROCS(2)
	fmt.Println(CalcularePi2(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
