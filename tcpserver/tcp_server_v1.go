package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	InitTcpServerV1()
}

func InitTcpServerV1() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError1(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError1(err, "Resolving address:port failed: "+hostAndPort)
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError1(err, "ListenTCP: ")
	fmt.Println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	connForm := conn.RemoteAddr().String()
	fmt.Println("Connection from: ", connForm)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	fmt.Println("Closed connection: ", connForm)
	checkError1(err, "Closed: ")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		fmt.Print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		fmt.Print(">")
	}
}

func sayHello(conn net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	write, err := conn.Write(obuf)
	checkError1(err, "Write: wrote "+fmt.Sprint(write)+" bytes.")
}

func checkError1(err error, str string) {
	if err != nil {
		panic("Error: " + str + err.Error())
	}
}
