package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ForRangeChannel1()
	ForRangeChannel2()
	time.Sleep(time.Second)
}

func ForRangeChannel1() {
	mychan := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			mychan <- strconv.Itoa(i)
		}
		close(mychan)
	}()
	go func() {
		for mychanitem := range mychan {
			fmt.Println("读取mychan：", mychanitem)
		}
	}()
}

func ForRangeChannel2() {
	mychan := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			mychan <- strconv.Itoa(i)
		}
		close(mychan)
	}()

	for {
		if mychanitem, ok := <-mychan; ok {
			fmt.Println("再次读取mychan：", mychanitem)
		} else {
			break
		}
	}
}
