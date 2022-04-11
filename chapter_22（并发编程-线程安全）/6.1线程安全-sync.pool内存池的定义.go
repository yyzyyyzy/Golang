package main

import (
	"fmt"
	"sync"
)

func main() {
	var p = &sync.Pool{
		New: func() interface{} {
			return "你好 golang1"
		},
	}
	p.Put("你好 golang2")
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
