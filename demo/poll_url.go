package demo

import (
	"fmt"
	"net/http"
)

func InitPollUrl() {
	urls := []string{
		"http://www.baidu.com/",
		// "http://golang.org/",
		// "http://blog.golang.org/",
	}

	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
