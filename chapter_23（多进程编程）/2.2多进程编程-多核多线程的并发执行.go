package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go AsyncFunc(i)
	}
	time.Sleep(time.Second * 1)
}

func AsyncFunc(index int) {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += 1
	}
	fmt.Printf("线程%d，sum为%d\n", index, sum)
}
