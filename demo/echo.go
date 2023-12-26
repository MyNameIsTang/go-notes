package demo

import (
	"flag"
	"os"
)

var newline = flag.Bool("n", false, "print newline")

const (
	space    = " "
	newline2 = "\n"
)

func InitEcho() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *newline {
				s += newline2
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
