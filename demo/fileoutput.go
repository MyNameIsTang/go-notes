package demo

import (
	"bufio"
	"fmt"
	"os"
)

func InitFileoutput() {
	outputFile, outputError := os.OpenFile("/Users/daguang/main/go-notes/demo/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "哈哈哈哈哈哈哈\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
