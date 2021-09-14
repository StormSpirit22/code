package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	log.Println("Start listening port 1234....")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func() {
			// 和之前唯一的不同就是 rpc.ServeConn 改成 rpc.ServeCodec 了。
			// 传入的参数是针对服务端的 json 编解码器
			// ServeCodec is like ServeConn but uses the specified codec to
			// decode requests and encode responses.
			rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		}()
	}
}
