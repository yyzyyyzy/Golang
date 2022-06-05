package main

import (
	"fmt"
	"time"
)

// 线程不安全的slice
var arr []int

func appendValue(i int) {
	arr = append(arr, i)
}

func main() {
	for i := 0; i < 10000; i++ {
		go appendValue(i)
	}

	go func() {
		for i, v := range arr {
			fmt.Printf("%d : %d\n", i, v)
		}
	}()
	time.Sleep(time.Second * 2)
}

// 线程不安全的map
func main1() {
	m := make(map[int]int)
	go func() {
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()

	time.Sleep(time.Second * 2)
}
