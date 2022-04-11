package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 10)
	timerchan := timer.C //chan Time.time
	fmt.Println("当前时间：", time.Now())
	fmt.Println("定时器时间：", <-timerchan)
}
