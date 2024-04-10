package main

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	InitWebsocketClient()
}

func InitWebsocketClient() {
	ws, err := websocket.Dial("ws://localhost:12345/websocket", "", "http://localhost/")
	if err != nil {
		panic("Dial: " + err.Error())
	}
	go readFromServer(ws)
	time.Sleep(5e9)
	ws.Close()
}

func readFromServer(ws *websocket.Conn) {
	ws.Write([]byte("哈哈哈哈哈"))
	buf := make([]byte, 1000)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}
}
