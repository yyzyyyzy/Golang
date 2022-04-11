package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	mychan := make(chan string, 3)

	go goroutine1(mychan)
	go goroutine2(mychan)
	go goroutine3(mychan)

	for i := 0; i < 3; i++ {
		fmt.Println(<-mychan)
	}

	time.Sleep(time.Second)
}
func goroutine1(mychan chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我日你仙人" + strconv.Itoa(i))
	}
	mychan <- "goroutine1 off"
}

func goroutine2(mychan chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我的天老爷" + strconv.Itoa(i))
	}
	mychan <- "goroutine2 off"
}

func goroutine3(mychan chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我的刘姥姥哎" + strconv.Itoa(i))
	}
	mychan <- "goroutine3 off"
}
