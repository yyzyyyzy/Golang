package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8848")
	HandleUDPClientErr(err, "net.Dial")

	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)
	for {
		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)
		n, _ := conn.Read(buffer)
		serverMsg := string(buffer[:n])
		fmt.Println("服务端：", serverMsg)

		if serverMsg == "im off" {
			conn.Write([]byte("服务端通知我溜了，那么我客户端也准备溜了")) //发送后，服务端返回消息，如果客户端已经退出，则会发生错误
			break
		}
	}
	fmt.Println("UDP Finished")

}

func HandleUDPClientErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
