package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var rwmutex sync.RWMutex
var wait sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wait.Add(1)
		go func() {
			rwmutex.RLock()
			fmt.Println("读取数据库") //锁定为只读模式，多路读取
			<-time.After(time.Second * 1)
			rwmutex.RUnlock()
			wait.Done()
		}()

		wait.Add(1)
		go func() {
			rwmutex.Lock()
			fmt.Println("写入数据库") //锁定为只写模式，单路写入
			<-time.After(time.Second * 1)
			rwmutex.Unlock()
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("程序结束")
}
