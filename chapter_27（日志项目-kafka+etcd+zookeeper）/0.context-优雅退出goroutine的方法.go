package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wait.Add(1)
	go worker(ctx)              //子协程
	time.Sleep(time.Second * 5) //主协程
	cancel()
	wait.Wait()
	fmt.Println("over")
}

func worker(ctx context.Context) {
	defer wait.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("我日你妈")
			time.Sleep(time.Second)
		}
	}
}
