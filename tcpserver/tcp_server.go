package tcpserver

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var data map[string]int

func main() {
	data = make(map[string]int)
	fmt.Println("Starting the server ...")
	listener, err := net.Listen("tcp", "localhost:50000")
	// if err != nil {
	// 	fmt.Println("Error listening", err.Error())
	// 	return
	// }
	checkError(err)
	for {
		conn, err := listener.Accept()
		checkError(err)
		// if err != nil {
		// 	fmt.Println("Error Accepting", err.Error())
		// 	return
		// }
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		checkError(err)
		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("server shutting down.")
			os.Exit(0)
		}

		if strings.Contains(input, ": WHO") {
			displayList()
		}

		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		data[clName] = 1
		fmt.Printf("Received data: --%v--", string(buf))
		// if err != nil {
		// 	fmt.Println("Error reading", err.Error())
		// 	return
		// }

		// fmt.Printf("Received data: %v \n", string(buf[:len]))
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error dialing" + err.Error())
	}
}

func displayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range data {
		fmt.Printf("user %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}
