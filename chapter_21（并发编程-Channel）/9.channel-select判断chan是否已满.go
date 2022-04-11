package main

import (
	"fmt"
	"time"
)

func main() {
	mychan := make(chan int, 10)

	go test(mychan)

	for {
		select {
		case mychan <- 666: //只能继续压入9个，已有一个
			fmt.Println("压入", 666)
		default:
			fmt.Println("mychan is full")
			return
		}
		time.Sleep(time.Second)
	}
}

func test(mychan chan int) {
	mychan <- 666
	time.Sleep(time.Second * 3)
}
