package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	var mutex sync.Mutex
	var money = 2000          //初始存款
	for i := 0; i < 10; i++ { //10人发红包
		wait.Add(1)
		go func() {
			mutex.Lock()                //操作数据前抢锁
			for j := 0; j < 1000; j++ { //每人发1000个1元红包
				money += 1
			}
			mutex.Unlock() //操作数据结束解锁
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(money) //应该为12000元
}
