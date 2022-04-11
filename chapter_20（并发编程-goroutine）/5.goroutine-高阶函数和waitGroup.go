package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	wait3.Add(1)
	toUpperSync("hello golang", func(s string) {
		toUpperSync(fmt.Sprintf("Result:%s\n", s), func(s string) { //无论如何嵌套都可以正常输出
			fmt.Printf("Result within:%s\n", s)
			wait3.Done()
		})

	})
	wait3.Wait()
}

var wait3 = new(sync.WaitGroup)

//小写转大写函数
func toUpperSync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()

}
