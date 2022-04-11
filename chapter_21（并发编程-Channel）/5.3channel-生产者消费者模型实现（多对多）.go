package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
生产者一直向【商店管道】写入
消费者一直向【商店管道】读出
消费者每读出一次，向【计数管道】写入一次
*/

type Product struct {
	name  int
	value int
}

func main() {
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	stop := false
	products := make(chan Product, 10)

	// 创建 5 个生产者和 5 个消费者
	for i := 1; i <= 5; i++ {
		go producer(&wgp, products, i, &stop)
		go consumer(&wgc, products, i)
		wgp.Add(1) //等待生产者协程
		wgc.Add(1) //等待消费者协程
	}

	time.Sleep(time.Duration(10) * time.Second)
	stop = true     // 设置生产者终止信号
	wgp.Wait()      // 等待生产者退出
	close(products) // 关闭通道
	wgc.Wait()      // 等待消费者退出
}

//如果 stop 标志不为 false，不断地往通道里面放 product，完成之后信号量完成
func producer(wg *sync.WaitGroup, products chan<- Product, name int, stop *bool) {
	for !*stop {
		product := Product{name: name, value: rand.Int()}
		products <- product
		fmt.Printf("producer %v produce a product: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

//不断地从通道里面取 product，然后作对应的处理，直到通道被关闭，并且 products 里面为空，for 循环才会终止，而这正是我们期望的
func consumer(wg *sync.WaitGroup, products <-chan Product, name int) {
	for product := range products {
		fmt.Printf("consumer %v consume a product: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}
