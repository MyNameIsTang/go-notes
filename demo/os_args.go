package demo

import (
	"fmt"
	"os"
	"strings"
)

func InitOsArgs() {
	who := "Hello "
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[2:], " ")
	}
	fmt.Printf("%s!", who)
}
