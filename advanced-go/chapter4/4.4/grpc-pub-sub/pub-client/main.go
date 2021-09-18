package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"rpc_protobuf/4.4/grpc-pub-sub/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &proto.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &proto.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
