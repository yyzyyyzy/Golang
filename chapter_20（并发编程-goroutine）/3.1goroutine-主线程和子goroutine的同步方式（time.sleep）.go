package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子协程" + strconv.Itoa(i))
		}
	}()
	fmt.Println("主线程退出")
	time.Sleep(1 * time.Second)
}
