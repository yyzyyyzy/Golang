package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	mychan := make(chan string, 5)
	for i := 0; i < 100; i++ {
		mychan <- "协程" + strconv.Itoa(i)
		go func() {
			time.Sleep(time.Second)
			fmt.Println(<-mychan)
		}()
	}
	time.Sleep(time.Second * 5)
}
