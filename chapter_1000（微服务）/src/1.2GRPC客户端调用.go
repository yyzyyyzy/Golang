package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	helloGrpc "protofile1"
)

func main() {
	// 客户端连接服务器
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// 新建一个客户端，获得grpc句柄
	client := helloGrpc.NewGreeterClient(conn)

	// 通过grpc句柄调用服务端函数
	r1, err := client.SayHello(context.Background(), &helloGrpc.HelloRequest{Name: "LZK"})
	if err != nil {
		fmt.Println("调用服务端代码失败1", err)
		return
	}
	fmt.Println("微服务1调用成功", r1.Message)

	r2, err := client.SayName(context.Background(), &helloGrpc.NameRequest{Name: "WHC"})
	if err != nil {
		fmt.Println("调用服务端代码失败2", err)
		return
	}
	fmt.Println("微服务2调用成功", r2.Message)

}
