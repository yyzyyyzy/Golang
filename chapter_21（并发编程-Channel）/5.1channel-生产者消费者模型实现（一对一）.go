package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	mychan := make(chan string)
	go Prodecer1(mychan)
	Consumer1(mychan)
	time.Sleep(time.Second)
}

func Prodecer1(mychan chan<- string) {
	for i := 0; i < 10; i++ {
		mychan <- "商品" + strconv.Itoa(i)
	}
	close(mychan)
}

func Consumer1(mychan <-chan string) {
	for str := range mychan {
		fmt.Println("消费的是：", str)
	}
}
