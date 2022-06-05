package client_proxy

import (
	"handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// Hello 封装client.Call方法
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}

// NewHelloServiceStub 封装rpc.Dial(network, address string) (*Client, error) {...}方法
func NewHelloServiceStub(network, address string) HelloServiceStub {
	conn, err := rpc.Dial(network, address)
	if err != nil {
		panic("rpc.Dial error!")
	}
	return HelloServiceStub{conn}
}
