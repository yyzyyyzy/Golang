package main

import (
	"fmt"
	"time"
)

func main() {
	channelStruct()
	dealNobufferdeadlock()
	time.Sleep(time.Second)
}

func channelStruct() {
	mychan1 := make(chan int)
	fmt.Println("管道的元素个数：", len(mychan1))
	fmt.Println("管道的总容量：", cap(mychan1))
}

func dealNobufferdeadlock() {
	//无缓冲的通道必须要有接受才能发送（同步通信），且发送和接收不能在相同的协程下执行，不然会引起死锁

	mychan2 := make(chan int)

	go func() {
		mychan2 <- 666
		close(mychan2)
	}()

	fmt.Println(<-mychan2)

}
