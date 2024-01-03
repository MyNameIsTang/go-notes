package demo

import (
	"fmt"
	"log"
)

func divide(dividend, divisor int) {
	fmt.Println("result:", dividend/divisor)
}

func divideWrap() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("出错啦！", err)
		}
	}()
	divide(10, 0)
}

func InitRecoverDividebyzero() {
	divideWrap()
}
