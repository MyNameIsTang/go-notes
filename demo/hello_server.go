package demo

import (
	"fmt"
	"net/http"
)

type CustomName struct {
	name string
}

func (c *CustomName) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprint(w, c.name)
}

func InitHelloServer() {
	cn := &CustomName{
		name: "Tom",
	}
	http.Handle("/", cn)
	http.ListenAndServe("localhost:9999", nil)
}
