package demo

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Counter struct {
	n int
}

type Chan chan int

var helloRequests = expvar.NewInt("hello-requests")
var webroot = flag.String("root", "/Users/daguang/main/go-notes", "web root directory")

func InitElaboratedWebServer() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(loggerfunc))
	http.Handle("/go/hello", http.HandlerFunc(HelloServer1))
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags", http.HandlerFunc(FlagServer))
	http.Handle("/args", http.HandlerFunc(ArgServer))
	http.Handle("/chan", ChanCreate())
	http.Handle("/date", http.HandlerFunc(DateServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}

func loggerfunc(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

func HelloServer1(w http.ResponseWriter, r *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello world!\n")
}

func FlagServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plan; charset=utf-8")
	fmt.Fprint(w, "Flags:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

func ArgServer(w http.ResponseWriter, r *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

func DateServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rx, wx, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(w, "pipe: %s\n", err)
		return
	}
	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, wx, wx}})
	defer rx.Close()
	wx.Close()
	if err != nil {
		fmt.Fprintf(w, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(w, rx)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(w, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(w, "date: %v\n", wait)
		return
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

func (c Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-c))
}

func (ctr *Counter) String() string {
	return fmt.Sprintf("%d", ctr.n)
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctr.n++
	case "POST":
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter =%d\n", ctr.n)
}
