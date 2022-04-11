package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 10)

	go func() {
		timerchan := timer.C
		fmt.Println("定时器时间：", <-timerchan)
	}()

	time.Sleep(time.Second * 5)
	stop := timer.Stop()
	if stop {
		fmt.Println("定时器停止成功")
	} else {
		fmt.Println("定时器停止失败")
	}
}
