package demo

import (
	"fmt"
	"net"
	"os"
)

func InitDial() {
	conn, err := net.Dial("tcp", "192.0.32.10:80")
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "192.0.32.10:80")
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80")
	checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is mode with %v\n", conn)
}
