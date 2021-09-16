package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"rpc_protobuf/4.3/rpc_watch/kvstore"
)

/*
server 运行 main.go 注册 kvstore service，会初始化 map 存储，一直跑不要断掉。
client 运行 main.go，将 abc 的值修改 set，就会输出 watch: abc
 */
func main() {
	service := kvstore.NewKVStoreService()
	rpc.RegisterName("KVStoreService", service)

	log.Println("Start listening port 1234....")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		fmt.Println("start conn ")

		// 这里会等到 conn 断掉再继续下一个连接
		go rpc.ServeConn(conn)
	}
}
