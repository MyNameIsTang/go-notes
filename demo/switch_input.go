package demo

import (
	"bufio"
	"fmt"
	"os"
)

func InitSwitchInput() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Printf("Your name is %s", input)

	switch input {
	case "Tom\n":
		fmt.Println("welcome Tom!")
	case "Jack\n":
		fmt.Println("welcome Jack!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	switch input {
	case "Tom\n":
		fallthrough
	case "Jack\n":
		fallthrough
	case "Chris\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	switch input {
	case "Tom\n", "Jack\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}
