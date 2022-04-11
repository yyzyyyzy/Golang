package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

func main() {
	b := Boy{User{1, "LZK", 20}, "NanJing"}

	fmt.Printf("%#v\n", reflect.TypeOf(b).Field(0)) // Anonymous：匿名

	fmt.Printf("%#v\n", reflect.ValueOf(b).Field(0)) // 值信息
}
