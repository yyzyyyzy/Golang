package main

import (
	"fmt"
	"time"
)

func main() {
	go runningTime()
	var input string
	fmt.Scan(&input)
}

func runningTime() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}
