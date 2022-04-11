package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	//拨号远程地址，建立TCP连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	ClientHandleError(err, "net.Dial")

	//开辟消息缓冲区
	buffer := make([]byte, 1024)
	//开辟命令行输入读取器（*File文件类型）
	reader := bufio.NewReader(os.Stdin)

	//一个完整的消息回合
	for {
		//接收命令行输入（一行消息）
		linebytes, _, _ := reader.ReadLine()

		//向服务端发送消息
		conn.Write(linebytes)

		//接收服务端消息，存入消息缓冲区，长度为n字节，
		n, _ := conn.Read(buffer)

		//转换为字节打印
		serverMsg := string(buffer[:n])
		fmt.Println("server:", serverMsg)

		//如果服务端发送断开请求，客户端也退出消息循环
		if serverMsg == "bye bye client" {
			break
		}
	}

	//客户端程序结束
	fmt.Println("TCP Connection Off")
}

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
