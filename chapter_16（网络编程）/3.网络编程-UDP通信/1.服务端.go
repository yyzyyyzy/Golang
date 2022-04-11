package main

import (
	"fmt"
	"net"
	"os"
)

func HandleUDPServerError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	//解析udp客户端的地址
	udp_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8848")
	HandleUDPServerError(err, "net.ResolverUDPAddr")

	//建立udp服务端监听，得到广口连接
	conn, err := net.ListenUDP("udp", udp_addr)
	HandleUDPServerError(err, "net.ListenUDP")
	defer conn.Close()

	//建立缓冲区
	buffer := make([]byte, 1024)

	//从广口连接中读入客户端发来的数据包，放入缓冲区
	n, remoteAddr, err := conn.ReadFromUDP(buffer) //读取一个数据包放到缓冲区，得到每个数据包的字节，远程地址
	HandleUDPServerError(err, "ReadFromUDP")

	//打印数据包的消息
	fmt.Printf("读到来自%v的内容：%s\n", remoteAddr, string(buffer[:n]))

	//回复客户端发送的内容
	conn.WriteToUDP([]byte("已阅"), remoteAddr)

}
