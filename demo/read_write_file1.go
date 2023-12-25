package demo

import (
	"fmt"
	"io/ioutil"
	"os"
)

func InitReadWriteFile1() {
	buf, err := ioutil.ReadFile("/Users/daguang/main/go-notes/demo/person.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile("/Users/daguang/main/go-notes/demo/person_copy.go", buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
