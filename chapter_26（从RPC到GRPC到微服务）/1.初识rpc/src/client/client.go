package main

import (
	"client_proxy"
	"fmt"
	"log"
)

func main() {
	//1.建立连接
	client := client_proxy.NewHelloServiceStub("tcp", "localhost:8005")

	//2.调用
	var reply string
	err := client.Hello("康少爷", &reply)
	if err != nil {
		log.Fatal("client.Call:", err)
	}
	fmt.Println(reply)
}
