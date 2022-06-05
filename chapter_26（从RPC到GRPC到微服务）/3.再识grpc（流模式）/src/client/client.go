package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"protofile"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//服务端流模式
	client := protofile.NewGreeterClient(conn)
	rep, _ := client.GetStream(context.Background(), &protofile.StreamReqData{Data: "fuck you"})
	for {
		a, err := rep.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	//客户端流模式
	putS, _ := client.PutStream(context.Background())
	i := 0
	for {
		i++
		putS.Send(&protofile.StreamReqData{Data: fmt.Sprintf("fuck you 慕课网%d", i)})
		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}
}
