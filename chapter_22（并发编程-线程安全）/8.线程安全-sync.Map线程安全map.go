package main

import (
	"fmt"
	"sync"
)

func main() {
	m := new(sync.Map)
	go func() {
		for {
			m.Store("name", "王浩川") //增/改
		}
	}()

	go func() {
		for {
			m.Load("name") //取/查
		}
	}()

	go func() {
		m.Range(func(key, value interface{}) bool { //遍历
			fmt.Println(key, value)
			return true
		})
	}()

	m.Delete("name") //删

}
