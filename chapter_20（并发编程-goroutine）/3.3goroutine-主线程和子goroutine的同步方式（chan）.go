package main

import (
	"fmt"
)

func main() {
	var chanTest = make(chan int)
	var chanMain = make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			chanTest <- i
			fmt.Println("生产者写入数据", i)
		}
		close(chanTest)
	}()

	go func() {
		for v := range chanTest {
			fmt.Println("\t消费者读出数据", v)
		}

		chanMain <- 666
	}()

	go func() {
		for v := range chanTest {
			fmt.Println("\t\t消费者读出数据", v)
		}
		chanMain <- 666
	}()

	<-chanMain
	<-chanMain

}
