package main

import (
	"fmt"
	"sync"
)

// 交替打印100以内的奇偶数

var (
	wg sync.WaitGroup
	ch = make(chan struct{})
)

func main() {
	wg.Add(2)
	go printOdd()
	go printEven()
	wg.Wait()
}

func printOdd() {
	for i := 1; i <= 100; i += 2 {
		<-ch
		fmt.Printf("奇数协程: num = %d\n", i)
		ch <- struct{}{}
	}
	wg.Done()
}

func printEven() {
	for i := 2; i <= 100; i += 2 {
		ch <- struct{}{}
		<-ch
		fmt.Printf("偶数协程: num = %d\n", i)
	}
	wg.Done()
}
