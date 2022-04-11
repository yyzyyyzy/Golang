package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。
//同时，当前的goroutine也会在未来的某个时间点继续运行。

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子协程" + strconv.Itoa(i))
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("主线程" + strconv.Itoa(i))
		}
	}()
	time.Sleep(time.Second)
}
