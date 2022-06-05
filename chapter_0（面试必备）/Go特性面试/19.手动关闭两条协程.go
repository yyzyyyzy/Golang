package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan interface{}, 10)

	go PrintNum(ch)
	go WriteNum(ch)

	for {
		i, isquit := <-ch //信号量控制协程关闭
		if !isquit {
			fmt.Println("PrintNum closed!")
			break
		}
		fmt.Println(i)
	}

	time.Sleep(time.Second)

}

func WriteNum(ch chan interface{}) {
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(10) //写随机数
	}
	close(ch)
	fmt.Println("WriteNum closed")
}

func PrintNum(ch chan interface{}) {
	for value := range ch {
		//time.Sleep(time.Second)
		fmt.Println(value) //读随机数
	}
}
