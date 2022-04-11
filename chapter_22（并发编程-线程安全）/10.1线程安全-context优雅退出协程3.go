package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	go Task4(ctx, "李子康")

	time.Sleep(time.Second * 3)
	fmt.Println("cancel 之前") //context.WithDeadline、context.WithTimeout:超时的时候就取消
	cancel()
	time.Sleep(time.Second * 3)
}

func Task4(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%v协程退出了\n", name)
			time.Sleep(time.Second)
		default:
			fmt.Printf("我%v好牛逼\n", name)
			time.Sleep(time.Second)
		}
	}
}
