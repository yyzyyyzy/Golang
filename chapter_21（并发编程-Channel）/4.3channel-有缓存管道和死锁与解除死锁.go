package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	dealbufferchannel()
	time.Sleep(time.Second)
}

func dealbufferchannel() {
	mychan := make(chan string, 5) //容量不会增长，会以队列FIFO形式入队列，出队列（异步通信）
	go func() {
		for i := 0; i < 10; i++ {
			mychan <- strconv.Itoa(i)
		}
		close(mychan) //防止死锁
	}()

	for mychanitem := range mychan {
		<-mychan
		fmt.Println(mychanitem)
	}
}
