package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan int)
	fmt.Println("start")
	time.AfterFunc(time.Second*2, func() {
		fmt.Println("boom!!!")
		exit <- 666
	})

	<-exit
}
