package client

import (
	"code/advanced-go/chapter4/4.1/2_safey_rpc/server/server"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ server.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(server.HelloServiceName+".Hello", request, reply)
}

