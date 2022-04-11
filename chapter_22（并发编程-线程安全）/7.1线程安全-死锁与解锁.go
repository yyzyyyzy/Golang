package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock()         //1.访问带锁的部分
		defer v1.mu.Unlock() //2.调用defer关键字释放锁

		time.Sleep(2 * time.Second) //3.添加休眠时间 以造成死锁
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b) //调用a锁定，然后尝试锁定b
	go printSum(&b, &a) //与此同时，锁定了b并尝试锁定a，两个goroutine都无限地等待着彼此
	wg.Wait()
}
