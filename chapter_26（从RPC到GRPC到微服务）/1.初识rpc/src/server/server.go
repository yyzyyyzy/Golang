package main

import (
	"handler"
	"log"
	"net"
	"net/rpc"
	"server_proxy"
)

func main() {
	//1.实例化server
	listener, err := net.Listen("tcp", ":8005") //监听套接字
	if err != nil {
		log.Fatal("listenTCP error :", err)
	}

	//2.注册微服务
	server_proxy.RegisterHelloService(&handler.HelloService{})

	//3.启动服务
	for {
		conn, err := listener.Accept() //新的请求进来，接收socket套接字
		if err != nil {
			log.Fatal("Accept error :", err)
		}
		rpc.ServeConn(conn)
	}
}
