package main

import "net"
import "fmt"

func main() {
	str := "1.2.3.4"
	address := net.ParseIP(str)
	if address != nil {
		fmt.Printf("%s is a legal ipv4 address\n", str)
	} else {
		fmt.Printf("%s is not a legal ipv4 address\n", str)
	}
}
