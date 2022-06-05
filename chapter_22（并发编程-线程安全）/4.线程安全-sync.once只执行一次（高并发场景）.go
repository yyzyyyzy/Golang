package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var data int64 = 0
var once sync.Once

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() { //无论多少次调用，只执行一次
				atomic.AddInt64(&data, 10)
			})
		}()
	}
	fmt.Println(data)
}
