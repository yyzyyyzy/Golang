package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	stoper := time.NewTimer(time.Second * 5)

	var i int
	for {
		select {
		case <-stoper.C:
			fmt.Println("超时")
			goto Quit
		case <-ticker.C:
			i++
			fmt.Println("任务" + strconv.Itoa(i))
		}
	}
Quit:
	fmt.Println("程序结束")
}
