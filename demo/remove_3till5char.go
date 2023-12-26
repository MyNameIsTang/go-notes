package demo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func InitRemove3Till5Char() {
	inputFile, _ := os.Open("/Users/daguang/main/go-notes/demo/products.txt")
	outputFile, _ := os.OpenFile("/Users/daguang/main/go-notes/demo/products_copy.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := string(inputString[2:5]) + "\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	outputWriter.Flush()
	fmt.Println("done")
}
