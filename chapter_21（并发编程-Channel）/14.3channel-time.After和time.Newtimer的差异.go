package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	time.After(time.Second * 2) //底层使用time.Newtimer实现的，此函数会马上返回，返回一个time.Time类型的Chan，不阻塞
	fmt.Println(reflect.TypeOf(time.After(time.Second * 2)))
	timeAfter()

	time.NewTimer(time.Second * 2) //可以实现循环计时，time.After只能配合select实现单次的超时停止，无法循环计时
	fmt.Println(reflect.TypeOf(time.NewTimer(time.Second * 2)))
	//Newtimer()
}

func Newtimer() {
	timer := time.NewTimer(time.Second * 2)
	defer timer.Stop()
	for {
		<-timer.C
		fmt.Println("timer running...")
		// 需要重置Reset 使 t 重新开始计时
		timer.Reset(time.Second * 2)
	}
}

func timeAfter() {
	mychan := make(chan int)
	isquit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-mychan:
				fmt.Println(num)
			case <-time.After(time.Second * 3):
				fmt.Println("超时")
				isquit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		mychan <- i
		time.Sleep(time.Second)
	}

	<-isquit
	fmt.Println("程序结束")
}
