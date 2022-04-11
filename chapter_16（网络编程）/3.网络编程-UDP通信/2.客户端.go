package main

import (
	"fmt"
	"net"
	"os"
)

func HandleUDPClientError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	//拨号请求连接
	conn, err := net.Dial("udp", "127.0.0.1:8848")
	HandleUDPClientError(err, "net.Dial")
	defer conn.Close()

	//发送消息给服务端
	n, err := conn.Write([]byte("客户端给服务端发送的第一条消息"))
	HandleUDPClientError(err, "conn.Write")
	fmt.Printf("消息占用字节数为%d\n", n)

	//创建缓冲区
	buffer := make([]byte, 1024)

	//从服务端中读入信息
	n, err = conn.Read(buffer)
	HandleUDPClientError(err, "conn.Read")
	fmt.Println(string(buffer[:n]))
}
