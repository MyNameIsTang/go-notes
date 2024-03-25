package demo

import (
	"fmt"
	"io"
	"net"
	"os"
)

func InitSocket() {
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]byte, 4096)
		read          = true
		count         = 0
	)
	con, err := net.Dial("tcp", remote)
	if err != nil {
		os.Exit(1)
	}
	io.WriteString(con, msg)
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Println(string(data[0:count]))
	}
	con.Close()
}
