package main

import (
	"fmt"
	"net"
	"os"
)

func HandleUDPServerErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8848")
	HandleUDPServerErr(err, "net.ResolveUDPAddr")

	udpConn, err := net.ListenUDP("udp", udp_addr)
	HandleUDPServerErr(err, "net.ListenUDP")

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, _ := udpConn.ReadFromUDP(buffer)
		clientMsg := string(buffer[:n])
		fmt.Printf("读到来自%v的内容：%s\n", remoteAddr, clientMsg)

		if clientMsg != "im off" {
			udpConn.WriteToUDP([]byte("已阅"), remoteAddr)
		} else {
			udpConn.WriteToUDP([]byte("bye bye"), remoteAddr)
		}

	}

}
