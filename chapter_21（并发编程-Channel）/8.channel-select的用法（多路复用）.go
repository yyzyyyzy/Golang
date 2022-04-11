package main

import (
	"fmt"
	"time"
)

// select从所有case中选择一条未阻塞的、最快的case执行，如果全部阻塞，那么执行default，或者等待一个可以继续通信的语句
func main() {
	output1 := make(chan int)
	output2 := make(chan int)
	go Test1(output1)
	go Test2(output2)

	for i := 0; i < 2; i++ {
		select { //多路复用：自由切换chan
		case n1 := <-output1:
			fmt.Println(n1)
		case n2 := <-output2:
			fmt.Println(n2)
		}
	}

}

func Test1(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1
}

func Test2(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 2
}
