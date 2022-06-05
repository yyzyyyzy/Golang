package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"protofile"
	"time"
)

type Server struct {
}

// GetStream 服务端流模式:客户端发起一次请求，服务端返回一段连续的数据流
func (s *Server) GetStream(req *protofile.StreamReqData, rep protofile.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = rep.Send(&protofile.StreamRepData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}
	return nil
}

// PutStream 客户端流模式:客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应
func (s *Server) PutStream(cliStr protofile.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

// AllStream 双向流模式:客户端和服务端都可以向对方发送数据流，这个时候双方的数据可以同时互相发送，也就是可以实现实时交互。
func (s *Server) AllStream(allStr protofile.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen " + err.Error())
	}

	g := grpc.NewServer()
	protofile.RegisterGreeterServer(g, &Server{})

	err = g.Serve(listener)
	if err != nil {
		panic(err)
	}
}
