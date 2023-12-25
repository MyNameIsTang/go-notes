package demo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type field struct {
	title string
	price float64
	count int
}

func InitReadCsv() {
	filename, _ := filepath.Abs("demo/products.txt")
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)
	arr := make([]field, 0)
	for {
		input, _, err := fileReader.ReadLine()
		re := strings.Split(string(input), ";")
		fmt.Println(re, len(re))
		if len(re) > 1 {
			price, _ := strconv.ParseFloat(re[1], 32)
			count, _ := strconv.Atoi(re[2])
			item := &field{
				re[0],
				price,
				count,
			}
			arr = append(arr, *item)
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println(arr)
}
