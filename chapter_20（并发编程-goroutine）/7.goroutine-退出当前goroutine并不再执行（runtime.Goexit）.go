package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。
//runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。
//请注意千万别在主函数调用runtime.Goexit，因为会引发panic。

func main() {
	go P1()
	go P2()
	fmt.Println(runtime.GOOS)
	time.Sleep(1 * time.Second)
}

func P1() {
	for i := 0; i < 5; i++ {
		fmt.Println("我是协程" + strconv.Itoa(i))
	}
}

func P2() {
	defer fmt.Println("知道啦1")
	runtime.Goexit()
}
