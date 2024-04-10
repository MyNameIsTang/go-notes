package tcpserver

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func InitRpcServer() {
	calc := new(Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Starting RPC-server - listen error: ", err)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
