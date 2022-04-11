package main

import (
	"log"
	"net"
	"net/rpc"
)

const FuckingServiceName = "FuckingService" //服务名

type FuckingService struct{}

func (f FuckingService) Fuck(penis string, vegina *string) error { //FuckingService结构体实现的Fuck方法
	*vegina = penis + "啊啊啊啊"
	return nil
}

type FuckingServiceInterface interface { //接口继承FuckingService结构体的Fuck方法
	Fuck(penis string, vegina *string) error
}

func RegisterFuckingService(svc FuckingServiceInterface) error { //封装rpc.RegisterName()
	return rpc.RegisterName(FuckingServiceName, svc)
}

func main() {
	RegisterFuckingService(new(FuckingService))
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}
