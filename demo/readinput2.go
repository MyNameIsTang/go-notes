package demo

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input1 string
var err error

func InitReadinput2() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input1, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input1)
	}
}
