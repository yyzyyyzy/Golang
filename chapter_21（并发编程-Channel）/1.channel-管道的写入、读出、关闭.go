package main

import (
	"fmt"
	"time"
)

func main() {
	CSPModel()
	time.Sleep(time.Second * 5)
}

func CSPModel() {
	mychan := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			mychan <- i //写入管道
			fmt.Println("写入管道的值：", i)
			time.Sleep(time.Second)
		}
		close(mychan) //关闭管道：再也无法写入数据，但是可以读取数据，关闭的管道读出的是数据类型的零值，通知读取的协程停止阻塞
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("读出管道的值：", <-mychan) //读出管道
		}
	}()
}
