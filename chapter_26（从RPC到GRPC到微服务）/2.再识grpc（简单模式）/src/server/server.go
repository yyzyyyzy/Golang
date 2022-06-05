package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"protofile"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *protofile.HelloRequest) (*protofile.HelloReply, error) {
	return &protofile.HelloReply{
		Message: "fuck you " + request.Name,
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":8080") //监听地址
	if err != nil {
		panic("failed to listen" + err.Error())
	}

	g := grpc.NewServer()                         // 初始化grpc服务
	protofile.RegisterGreeterServer(g, &Server{}) // 注册服务

	err = g.Serve(listener) // 监听服务，如启动失败则抛出异常
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
}
