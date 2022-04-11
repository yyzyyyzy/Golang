package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	StartBingfafunction2 := func() {
		for i := 0; i < 10; i++ {
			fmt.Println("并发执行" + strconv.Itoa(i))
		}
	}
	//实现方法：
	go StartBingfafunction1() //go关键字实现goroutine
	go StartBingfafunction2() //命名函数实现goroutine
	go func() {               //匿名函数实现goroutine
		for i := 0; i < 10; i++ {
			fmt.Println("并发执行" + strconv.Itoa(i))
		}
	}()
	time.Sleep(time.Second)
}

func StartBingfafunction1() {
	for i := 0; i < 10; i++ {
		fmt.Println("并发执行" + strconv.Itoa(i))
	}
}
