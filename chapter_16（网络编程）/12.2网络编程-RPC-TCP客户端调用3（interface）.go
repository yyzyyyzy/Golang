package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const FuckingServiceName = "FuckingService" //服务名

type FuckingServiceInterface interface {
	Fuck(penis string, vegina *string) error
}

type FuckingServiceClient struct {
	*rpc.Client //客户端
}

var _ FuckingServiceInterface = (*FuckingServiceClient)(nil) //匿名遍历

// 客户端调用实际方法
func (f FuckingServiceClient) Fuck(penis string, vegina *string) error {
	return f.Client.Call(FuckingServiceName+".Fuck", penis, vegina)
}

// 客户端远程调用
func DialFuckingService(network, address string) (*FuckingServiceClient, error) {
	conn, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &FuckingServiceClient{conn}, nil
}

func main() {
	client, err := DialFuckingService("tcp", "localhost:8083")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Fuck("黑山雀", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
