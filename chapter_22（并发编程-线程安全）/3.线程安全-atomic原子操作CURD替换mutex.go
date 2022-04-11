package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//用原子操作来替换mutex锁的主要原因是：
//原子操作由底层硬件支持，而锁则由操作系统提供的API实现。若实现相同的功能，前者通常会更有效率。

func main() {
	var num int64 = 50

	for i := 0; i < 10; i++ {
		wait1.Add(1)
		go func() {
			Add(&num)
			wait1.Done()
		}()
	}
	wait1.Wait()
	fmt.Println("num=", num)
	fmt.Println("程序结束")
}

var wait1 sync.WaitGroup

func Add(num *int64) {
	atomic.AddInt64(num, 100000)
}
