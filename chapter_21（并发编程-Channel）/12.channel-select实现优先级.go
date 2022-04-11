package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	mychan1 := make(chan int, 10)
	mychan2 := make(chan int, 10)
	isquit := make(chan bool)
	go test1(mychan1, mychan2, isquit)
	go func() {
		for i := 0; i < 10; i++ {
			mychan1 <- 666
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mychan2 <- i
		}
	}()
	time.Sleep(time.Second)

}

func test1(mychan1, mychan2 chan int, isquit chan bool) {
	for {
		select {
		case <-isquit:
			runtime.Goexit()
		case task1 := <-mychan1: //优先处理task1
			fmt.Println(task1)
		case task2 := <-mychan2:
		priority: //label标签
			for {
				select {
				case task1 := <-mychan1:
					fmt.Println(task1)
				default:
					break priority
				}
			}
			fmt.Println(task2)
		}
	}
}
