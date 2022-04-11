package main

import (
	"fmt"
)

func main() {
	mychan := make(chan int)
	isquit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			mychan <- i
		}
		isquit <- true
	}()

	GetNum(mychan, isquit)
}
func GetNum(mychan chan int, isquit chan bool) {
	for {
		select {
		case num := <-mychan:
			fmt.Println(num)
		case <-isquit:
			return
		}
	}
}
