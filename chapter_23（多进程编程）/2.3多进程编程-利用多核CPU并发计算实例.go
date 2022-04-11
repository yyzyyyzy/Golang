package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	for i := 1; i <= 10000000; i++ {
		sum += i
	}
	fmt.Printf("子协程%d运算结果:%d\n", seq, sum)
	ch <- sum
}

func main() {
	// 启动时间
	start := time.Now()
	// 最大 CPU 核心数
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	chs := make([]chan int, cpus)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, cpus)
		go sum(i, chs[i])
	}
	sum := 0
	for _, ch := range chs {
		res := <-ch
		sum += res
	}
	// 结束时间
	end := time.Now()
	// 打印耗时
	fmt.Printf("最终运算结果: %d, 执行耗时(s): %f\n", sum, end.Sub(start).Seconds())
}
