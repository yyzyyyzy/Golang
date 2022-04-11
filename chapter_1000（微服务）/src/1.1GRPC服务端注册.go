package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	helloGrpc "protofile1"
)

type server struct {
}

// SayHello 服务1
func (s *server) SayHello(ctx context.Context, request *helloGrpc.HelloRequest) (*helloGrpc.HelloReply, error) {
	return &helloGrpc.HelloReply{Message: "hello" + request.Name}, nil
}

// SayName 服务2
func (s *server) SayName(ctx context.Context, request *helloGrpc.NameRequest) (*helloGrpc.NameReply, error) {
	return &helloGrpc.NameReply{Message: request.Name}, nil
}

func main() {
	// 监听本地端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
		return
	}

	// 创建gRPC服务器
	newserver := grpc.NewServer()
	helloGrpc.RegisterGreeterServer(newserver, &server{})
	reflection.Register(newserver)

	err = newserver.Serve(listen)
	if err != nil {
		fmt.Println("开启服务失败", err)
		return
	}
}
