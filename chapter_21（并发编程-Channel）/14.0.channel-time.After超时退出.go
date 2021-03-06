package main

import (
	"fmt"
	"time"
)

func main() {
	mychan := make(chan int)
	isquit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-mychan: //如果有数据，下面打印。但是有可能ch一直没数据
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second): //上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
				fmt.Println("超时")
				isquit <- true //写入
			}
		}

	}()

	for i := 1; i <= 5; i++ {
		mychan <- i
		time.Sleep(time.Second)
	}

	<-isquit //这里暂时阻塞，直到可读
	fmt.Println("程序结束")

}
