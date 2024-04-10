package tcpclient

import (
	"basic/tcpserver"
	"fmt"
	"log"
	"net/rpc"
)

func InitRpcClient() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error dialing: ", err)
	}
	args := &tcpserver.Args{N: 7, M: 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
