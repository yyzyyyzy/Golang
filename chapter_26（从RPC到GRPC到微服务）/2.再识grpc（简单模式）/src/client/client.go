package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"protofile"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) // 建立链接
	if err != nil {
		panic("failed to Dial" + err.Error())
	}

	defer conn.Close()

	client := protofile.NewGreeterClient(conn)                                                                          // 初始化客户端
	reply, err := client.SayHello(context.Background(), &protofile.HelloRequest{Name: "LZK", Age: 18, Addr: "Nanjing"}) // 调用SayHello rpc服务
	if err != nil {
		panic("failed to SayHello" + err.Error())
	}
	fmt.Println(reply.Message)
}
