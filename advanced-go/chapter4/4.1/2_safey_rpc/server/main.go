package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_protobuf/4.1/2_safey_rpc/server/server"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = request + " world"
	return nil
}

func main() {
	server.RegisterHelloService(new(HelloService))

	log.Println("Start listening port 1234....")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 支持多个TCP链接，然后为每个TCP链接提供RPC服务。
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
