package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Task1(ctx, "李子康")
	go Task1(ctx, "王浩川")
	go Task1(ctx, "李泽熙")
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 3)
}

func Task1(ctx context.Context, name string) {
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
