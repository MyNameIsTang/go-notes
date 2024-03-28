package demo

import (
	"io"
	"net/http"
)

const form = `
	<html>
		<body>
			<form action="#" method="post" name="bar">
				<input type="text" name="in" />
				<input type="submit" value="submit"/>
			</form>
		</body>
	</html>
`

func InitSimpleWebServer() {
	http.HandleFunc("/test1", simpleServer)
	http.HandleFunc("/test2", formServer)
	err := http.ListenAndServe("localhost:8089", nil)
	checkError(err)
}

func simpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello world</h1>")
}

func formServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, req.FormValue("in"))
	}
}
