package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	messageGPC "protofile2"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := messageGPC.NewOrderServiceClient(conn)
	orderRequest := &messageGPC.OrderRequest{OrderId: "20220328", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
