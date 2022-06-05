package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func producer(ch chan<- interface{}) {
	for i := 0; i < 5; i++ {
		ch <- "生产的商品" + strconv.Itoa(rand.Intn(10))
	}
}

func consumer(ch <-chan interface{}) {
	for value := range ch {
		fmt.Println("消费者消费了", value)
	}
}

func main() {
	mychan := make(chan interface{}, 5)
	for i := 0; i < 10; i++ {
		go producer(mychan)
	}

	go consumer(mychan)

	time.Sleep(time.Second * 3)

}
