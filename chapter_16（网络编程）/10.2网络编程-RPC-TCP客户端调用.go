package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081") //客户端函数rpc.DialHTTP创建客户端句柄client
	if err != nil {
		panic(err.Error())
	}
	add := new(AddStruct)
	add.Num1 = 10
	add.Num2 = 30
	var result *int
	err = client.Call("AddStruct.AddFunc", add, &result) //客户端调用远程函数（对象.函数）同时传参和接收返回值
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*result)

}

type AddStruct struct {
	Num1 int
	Num2 int
}
