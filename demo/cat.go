package demo

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberline = flag.Bool("nl", false, "print newline")

func cat(r *bufio.Reader) {
	i := 0
	for {
		buf, err := r.ReadBytes('\n')
		if *numberline {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
		if err == io.EOF {
			break
		}
	}
}

func InitCat() {
	flag.PrintDefaults()
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
