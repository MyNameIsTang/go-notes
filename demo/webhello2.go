package demo

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func InitWebHello2() {
	http.HandleFunc("/hello/", helloName)
	http.HandleFunc("/shouthello/", helloName)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}

func helloName(w http.ResponseWriter, req *http.Request) {
	arr := strings.Split(req.URL.Path, "/")
	last := arr[len(arr)-1]
	fmt.Println(last)
	fmt.Fprintf(w, "hello %s", last)
}
