package main

import (
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	addStruct := new(AddStruct)

	//1.注册结构体
	err := rpc.Register(addStruct)
	if err != nil {
		panic(err.Error())
	}

	//2.将结构体提供的服务注册到http协议内
	rpc.HandleHTTP()

	//3.在指定的端口进行监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}

	http.Serve(listen, nil)
}

type AddStruct struct {
	Num1 int
	Num2 int
}

func (a AddStruct) AddFunc(add AddStruct, result *int) error {
	*result = add.Num1 + add.Num2
	return nil
}
