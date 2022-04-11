package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 条件变量解决的问题：生产者太多消费不过来，消费者太多生产不过来，所以需要条件变量和锁搭配使用实现负载均衡

func main() {
	ch := make(chan int, 3)          //定义管道，用于通信
	quit := make(chan bool)          //判断是否退出
	cond.L = new(sync.Mutex)         //定义条件变量用的锁
	rand.Seed(time.Now().UnixNano()) //随机数种子
	for i := 1; i < 15; i++ {
		go producer(ch, i)
	}
	for i := 1; i < 6; i++ {
		go consummer(ch, i)
	}
	<-quit

}

var cond sync.Cond //定义全局条件变量

func producer(out chan<- int, idx int) { //定义生产者
	for {
		cond.L.Lock()       //给公共缓冲区加锁
		for len(out) == 3 { //循环判断缓冲区容量,缓冲区满了，阻塞
			cond.Wait() //1）阻塞2）解锁3）加锁。 1和2是一个原子操作
		}
		num := rand.Intn(1000) //产生随机数
		out <- num             //随机数写入管道
		fmt.Printf("生产者:%d,生产:%d\n", idx, num)
		cond.L.Unlock()  //给公共缓冲区解锁
		cond.Broadcast() //唤醒被阻塞的消费者，进行消费数据
		time.Sleep(time.Millisecond * 300)
	}
}
func consummer(in <-chan int, idx int) { //定义消费者
	for {
		cond.L.Lock()      //给公共缓冲区加锁
		for len(in) == 0 { //循环判断缓冲区容量，是否为空，为空阻塞，不为空，继续消费数据
			cond.Wait() //1）阻塞2）解锁3）加锁。1和2是一个原子操作
		}
		num := <-in //从管道中消费随机数
		fmt.Printf("##########消费者:%d,消费:%d\n", idx, num)
		cond.L.Unlock()  //给公共缓冲区解锁
		cond.Broadcast() //管道中数据被消费没了，唤醒被阻塞的生产者，进行生产数据
		time.Sleep(time.Millisecond * 200)
	}
}
