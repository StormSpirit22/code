package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc_protobuf/4.4/grpc/proto"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *proto.String,
) (*proto.String, error) {
	reply := &proto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}