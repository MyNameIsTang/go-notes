package even

import "fmt"

func MainOddven(num int) {
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			fmt.Printf("the value is even: %v\n", i)
		}
	}
}

func Even(i int) bool {
	return i%2 == 0
}

func Odd(i int) bool {
	return i%2 != 0
}
