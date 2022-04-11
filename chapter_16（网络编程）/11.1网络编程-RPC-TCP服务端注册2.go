package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	//1.注册结构体
	rpc.RegisterName("HelloService", new(HelloService))

	//2.在指定端口进行监听
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("listenTCP error :", err)
	}

	//3.接收连接,for用来服务多个连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error :", err)
		}
		rpc.ServeConn(conn)
	}
}
