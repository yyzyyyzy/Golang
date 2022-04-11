package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func HandleTcpClientError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	HandleTcpClientError(err, "net.Dial")

	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		lineBytes, _, _ := reader.ReadLine()
		conn.Write(lineBytes)
		n, _ := conn.Read(buffer)
		serverMsg := string(buffer[:n])
		fmt.Println("server:", serverMsg)

		if serverMsg == "bye bye client" {
			break
		}
	}
	fmt.Println("TCP Connection is off")
}
