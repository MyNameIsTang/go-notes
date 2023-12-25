package demo

import (
	"bufio"
	"fmt"
	"os"
)

func InitCalculator() {
	stack := new(Stack2)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter number1 and number2 and operator:")

	func1 := func() int {
		operator, _ := stack.Pop()
		number2, _ := stack.Pop()
		number1, _ := stack.Pop()
		var a, b int

		if v, ok := number1.(int); ok {
			a = v
		}
		if v, ok := number2.(int); ok {
			b = v
		}

		switch operator {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		default:
			return 0
		}
	}

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There were errors reading, exiting program.")
			return
		}
		if input == "q\n" {
			fmt.Println(stack.arr)
			fmt.Printf("The result is %d\n", func1())
			os.Exit(0)
		}
		if stack.Len() < 2 {
			var i int
			fmt.Sscanf(input, "%d", &i)
			stack.Push(i)
		} else {
			var i string
			fmt.Sscanf(input, "%s", &i)
			stack.Push(i)
		}
	}
}
