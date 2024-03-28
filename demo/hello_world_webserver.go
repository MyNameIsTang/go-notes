package demo

import (
	"fmt"
	"log"
	"net/http"
)

func InitHelloWorldWebServer() {
	// http.HandleFunc("/", helloServer)
	http.Handle("/", http.HandlerFunc(helloServer))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside helloserver handler")
	fmt.Fprintf(w, "<h1>%s</h1>", req.URL.Path[1:])
	// http.Redirect(w, req, "/haha", 200)
	// http.NotFound(w, req)
	// http.Error(w, "我是来搞笑的", 500)
}
