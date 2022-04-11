package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	messageGPC "protofile2"
	"time"
)

type OrderServiceImpl struct {
}

func (ov *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *messageGPC.OrderRequest) (*messageGPC.OrderInfo, error) {
	orderMap := map[string]messageGPC.OrderInfo{
		"20220326": messageGPC.OrderInfo{OrderId: "20220326", OrderName: "衣服", OrderStatus: "已付款"},
		"20220327": messageGPC.OrderInfo{OrderId: "20220327", OrderName: "裤子", OrderStatus: "已付款"},
		"20220328": messageGPC.OrderInfo{OrderId: "20220328", OrderName: "手套", OrderStatus: "未付款"},
	}
	var response *messageGPC.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = messageGPC.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("服务器错误")
		}
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()
	messageGPC.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
