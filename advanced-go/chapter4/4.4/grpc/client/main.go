package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rpc_protobuf/4.4/grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}