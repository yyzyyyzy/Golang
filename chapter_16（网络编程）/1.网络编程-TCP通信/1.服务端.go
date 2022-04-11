package main

import (
	"fmt"
	"net"
	"os"
)

func HandleTcpServerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1) //code=1暴力退出且不再执行defer函数， code=0正常退出
	}
}

func ioWithConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		HandleTcpServerError(err, "conn.Read")
		clientMsg := string(buffer[:n])
		fmt.Printf("received client:%v ,msg:%s\n", conn.RemoteAddr(), clientMsg)

		if clientMsg == "im off" {
			conn.Write([]byte("bye bye client"))
			break
		}
		conn.Write([]byte("clientMsg received" + clientMsg))
	}
	fmt.Println("client is off")
}
func main() {
	listenerSocket, err := net.Listen("tcp", "127.0.0.1:8888")
	HandleTcpServerError(err, "net.Listen")
	fmt.Println("Listening....")

	for {
		conn, err := listenerSocket.Accept() //等待连接
		HandleTcpServerError(err, "Listen.Accept")
		ioWithConn(conn)
	}
}
