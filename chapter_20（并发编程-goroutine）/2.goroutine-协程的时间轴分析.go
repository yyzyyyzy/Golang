package main

import (
	"fmt"
	"strconv"
	"time"
)

//协程有初始化的时间，当子协程初始化完毕时，主协程已经运行很久了
//主线程time.sleep()退出，子协程也会退出
func main() {
	go PrintNum1()
	go PrintNum2()
	time.Sleep(10 * time.Second)
}

func PrintNum1() {
	for i := 0; i < 10; i++ {
		fmt.Println("并发任务A" + strconv.Itoa(i))
		time.Sleep(500 * time.Millisecond)
	}
}

func PrintNum2() {
	for i := 0; i < 10; i++ {
		fmt.Println("并发任务B" + strconv.Itoa(i))
		time.Sleep(250 * time.Millisecond)
	}
}
