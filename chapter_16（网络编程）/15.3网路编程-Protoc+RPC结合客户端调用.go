package main

import (
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	timestamp := time.Now().Unix()
	request := OrderRequest{OrderId: "20220328", TimeStamp: timestamp}

	var response *OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(*response)
}

// 从message.pb.go 拷贝的部分
type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	TimeStamp int64  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     string `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	OrderName   string `protobuf:"bytes,2,opt,name=OrderName,proto3" json:"OrderName,omitempty"`
	OrderStatus string `protobuf:"bytes,3,opt,name=OrderStatus,proto3" json:"OrderStatus,omitempty"`
}
