package demo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitWordLetterCount() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter something: ")
	// if err != nil {
	// 	fmt.Println("There were errors reading, exiting program.")
	// 	return
	// }
	var count1, count2, count3 int
	// for _, v := range input1 {
	// 	count2++
	// 	if v != '\r' && v != '\n' {
	// 		count1++
	// 	}
	// 	if v == '\n' {
	// 		count3++
	// 	}
	// }
	// if count3 != 0 {
	// 	count3++
	// }
	// fmt.Printf("输入的字符的个数:%d, 输入的单词的个数:%d,输入的行数:%d", count1, count2, count3)

	counters := func(input string) {
		count1 += len(input) - 2
		count2 += len(strings.Fields(input))
		count3++
	}

	for {
		input1, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There were errors reading, exiting program.")
			return
		}
		if input1 == "S\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", count1)
			fmt.Printf("Number of words: %d\n", count2)
			fmt.Printf("Number of lines: %d\n", count3)
			os.Exit(0)
		}
		counters(input1)
	}
}
