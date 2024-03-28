package demo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func InitHttpFetch() {
	fmt.Println("输入url")
	for {
		reader := bufio.NewReader(os.Stdin)
		content, err := reader.ReadString('\n')
		checkError(err)
		if content == "Q" {
			os.Exit(0)
		}

		content = strings.Trim(content, "\n")
		if content == "" {
			continue
		}
		res, err := http.Get(content)
		checkError(err)
		data, err := io.ReadAll(res.Body)
		checkError(err)
		fmt.Printf("Got: %q", string(data))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
