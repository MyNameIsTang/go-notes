package fibo

import "fmt"

var Fib [50]int

func Fibonacci(op string, i int) (res int) {
	if Fib[i] != 0 {
		res = Fib[i]
		return
	}

	if i <= 1 {
		res = 1
	} else {
		fmt.Println(i, op)
		v1 := Fibonacci(op, i-1)
		v2 := Fibonacci(op, i-2)
		if op == "*" {
			res = v1 * v2
		} else {
			res = v1 + v2
		}
	}
	fmt.Println(res)
	Fib[i] = res
	return
}
