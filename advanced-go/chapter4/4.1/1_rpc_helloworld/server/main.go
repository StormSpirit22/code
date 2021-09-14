package main

import (
	"log"
	"net"
	"net/rpc"
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

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	log.Println("Start listening port 1234....")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}