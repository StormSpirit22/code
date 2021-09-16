package main

import (
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {}

/*
Hello
如果对象的方法要能远程访问，它们必须满足一定的条件：
方法的类型是可输出的 (the method's type is exported)
方法本身也是可输出的 （the method is exported）
方法必须由两个参数，必须是输出类型或者是内建类型 (the method has two arguments, both exported or builtin types)
方法的第二个参数是指针类型 (the method's second argument is a pointer)
方法返回类型为 error (the method has return type error)
 */
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

// 反向RPC的内网服务将不再主动提供TCP监听服务，而是首先主动链接到对方的TCP服务器。然后基于每个建立的TCP链接向对方提供RPC服务。
func main() {
	rpc.Register(new(HelloService))

	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}