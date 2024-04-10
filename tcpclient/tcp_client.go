package tcpclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:50000")
// 	if err != nil {
// 		fmt.Println("Error dialing", err.Error())
// 		return
// 	}
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("First, what is your name?")
// 	clientName, _ := inputReader.ReadString('\n')
// 	trimmedClient := strings.Trim(clientName, "\n")
// 	for {
// 		fmt.Println("What to send to the server? Type Q to quit.")
// 		input, _ := inputReader.ReadString('\n')
// 		trimmedInput := strings.Trim(input, "\n")
// 		if trimmedInput == "Q" {
// 			return
// 		}
// 		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
// 	}
// }

func InitTcpClient() {
	conn, err := net.Dial("tcp", "localhost:50000")
	checkError(err)
	fmt.Println("First, what is your name?")
	inputReader := bufio.NewReader(os.Stdin)
	clientName, _ := inputReader.ReadString('\n')
	name := strings.Trim(clientName, "\n")
	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to stop server")
		input, _ := inputReader.ReadString('\n')
		trimedInput := strings.Trim(input, "\n")
		if trimedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(name + " says: " + trimedInput))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error dialing" + err.Error())
	}
}
