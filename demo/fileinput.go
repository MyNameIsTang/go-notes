package demo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func InitFileInput() {
	inputFile, inputError := os.Open("/Users/daguang/main/go-notes/demo/person.go")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
