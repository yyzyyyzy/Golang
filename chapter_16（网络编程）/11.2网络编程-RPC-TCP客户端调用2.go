package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8082")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
