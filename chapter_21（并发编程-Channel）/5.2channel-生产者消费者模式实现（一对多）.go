package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	mychan := make(chan string)
	for i := 0; i < 5; i++ {
		go Producer2(mychan)
	}
	go Consumer2(mychan)

	time.Sleep(time.Second)
}

func Producer2(mychan chan<- string) {
	for i := 0; i < 10; i++ {
		mychan <- "商品" + strconv.Itoa(i)
	}
}

func Consumer2(mychan <-chan string) {
	for str := range mychan {
		fmt.Println("消费的是：", str)
	}
}
