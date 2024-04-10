package tcpserver

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func server2(ws *websocket.Conn) {
	fmt.Printf("new connection\n")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
	}
	fmt.Println("value: ", string(buf))
	fmt.Printf(" => closing connection\n")
	ws.Close()
}

func InitWebsocketServer() {
	http.Handle("/websocket", websocket.Handler(server2))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
