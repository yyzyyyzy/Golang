package main

import (
	"fmt"
	"runtime"
)

func Fibonacci(mychan chan int, isquit chan bool) {
	for {
		select {
		case num := <-mychan: //能够取出，打印数据
			fmt.Println(num)
		case <-isquit: //能够取出，退出线程
			runtime.Goexit()
		}
	}
}
func main() {
	mychan := make(chan int)  //数据管道
	isquit := make(chan bool) //判断停止管道

	go Fibonacci(mychan, isquit)

	a, b := 1, 1
	for i := 0; i < 50; i++ {
		mychan <- a
		a, b = b, a+b
	}
	isquit <- true

}
