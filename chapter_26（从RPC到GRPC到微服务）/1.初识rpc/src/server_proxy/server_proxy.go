package server_proxy

import (
	"handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(name HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, name)
}
