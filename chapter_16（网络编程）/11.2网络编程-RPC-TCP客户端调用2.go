package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//1.建立连接
	client, err := rpc.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//2.调用
	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
