package main

import (
	"fmt"
	"time"
)

/* 你吃饭吃到一半，电话来了，你一直到吃完了以后才去接，这就说明你不支持并发也不支持并行。

你吃饭吃到一半，电话来了，你停了下来接了电话，接完后继续吃饭，这说明你支持并发。

你吃饭吃到一半，电话来了，你一边打电话一边吃饭，这说明你支持并行。*/

func EatAndPhone() {
	fmt.Println("吃饭中。。。。")
	time.Sleep(5 * time.Second)
	fmt.Println("接电话中。。。")
}

func Begin() {
	for i := 0; i < 10; i++ {
		EatAndPhone()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go EatAndPhone() //并发乱序执行
	}
	time.Sleep(6 * time.Second) //主线程，我们需要主go程后于子go程结束，防止主goroutine退出后，其它的工作goroutine也会自动退出的情况
}
