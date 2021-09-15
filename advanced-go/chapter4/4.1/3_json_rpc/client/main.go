package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 先手工调用net.Dial函数建立TCP链接，然后基于该链接建立针对客户端的json编解码器。
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial: ", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

/*
使用 tcp 服务而不是 Go 的 rpc 服务（即停掉 server)，命令行里输入 nc -l 1234
在 client 里运行 go run plugin.go
nc 命令行会输出：

{"method":"HelloService.Hello","params":["hello"],"id":0}

然后在 server 里运行 go run plugin.go，在命令行里输入

 echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234

返回 {"id":1,"result":"hello:hello","error":null}
 */