package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go RequestTask(ctx, "LZK")
	time.Sleep(time.Second * 3)
	cancel() //context.WithCancel():执行取消函数就取消
	time.Sleep(time.Second * 3)
}

func RequestTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("协程结束，退出")
			return
		default:
			fmt.Printf("协程%v正在运行\n", name)
			time.Sleep(time.Second)
		}
	}
}
