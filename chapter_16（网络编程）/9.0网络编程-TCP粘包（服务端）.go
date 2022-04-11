package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer con.Close()

		for {
			var data = make([]byte, 5)
			n, err := con.Read(data)
			if err != nil && err != io.EOF {
				log.Println(err)
			}
			if n > 0 {
				log.Println("received msg", n, "bytes:", string(data[:n]))
			}
		}
	}
}
