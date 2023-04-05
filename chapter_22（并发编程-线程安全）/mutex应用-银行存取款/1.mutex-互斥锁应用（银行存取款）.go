package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	var wait sync.WaitGroup

	account := new(Account)
	account.money = 3000

	for i := 0; i < 3; i++ {
		wait.Add(1)
		go func() {
			account.Get(1000)
			wait.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wait.Add(1)
		go func() {
			account.Save(3000)
			wait.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wait.Add(1)
		go func() {
			account.Query()
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("程序结束")
}

type Account struct {
	money int
}

// 取钱（需要并发安全，加锁）
func (a *Account) Get(n int) {
	mutex.Lock()
	fmt.Println("取钱 开始")
	<-time.After(time.Second * 3)
	a.money -= n
	fmt.Println("取钱 结束")
	mutex.Unlock()
}

// 存钱（需要并发安全，加锁）
func (a *Account) Save(n int) {
	mutex.Lock()
	fmt.Println("存钱 开始")
	<-time.After(time.Second * 3)
	a.money += n
	fmt.Println("存钱 结束")
	mutex.Unlock()
}

// 查询（无需并发安全）
func (a *Account) Query() {
	fmt.Println("查询 开始")
	<-time.After(time.Second * 3)
	fmt.Println("余额：", a.money)
	fmt.Println("查询 结束")
}
