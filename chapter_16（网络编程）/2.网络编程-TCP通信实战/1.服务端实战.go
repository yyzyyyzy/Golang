package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//服务端建立监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	ServerHandleError(err, "net.listen")
	fmt.Println("Listening....")

	for {
		//循环接入所有客户端，得到专线连接
		conn, err := listener.Accept() //写法是阻塞的
		ServerHandleError(err, "listen.Accept")
		//每个用户开启独立协程，与该客户端聊天
		go Process(conn)
	}
}

// 在Conn中与客户对话
func Process(conn net.Conn) {

	//创建消息缓冲区
	buffer := make([]byte, 1024)

	//一个完整的消息回合，不断的收发消息
	for {
		// 读取客户端发来的消息，存入缓冲区，消息长度为n字节
		n, err := conn.Read(buffer)
		ServerHandleError(err, "conn.Read")

		//转换为字符串输出
		clientMsg := string(buffer[:n]) //客户端发来的消息
		fmt.Printf("received client:%v ,msg:%s\n", conn.RemoteAddr(), clientMsg)

		if clientMsg == "TCP OFF" {
			conn.Write([]byte("bye bye client")) //服务端回应客户端
			break
		}
		//回复客户端消息
		conn.Write([]byte("client server" + clientMsg))
	}

	//断开客户端连接
	conn.Close()
	fmt.Printf("client%s is off\n", conn.RemoteAddr()) //conn.RemoteAddr()客户端远程地址
}

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
