package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err", err)
	}
	err = conn.Send("auth", "root")
	if err != nil {
		fmt.Println("conn.Send err", err)
	}
	defer conn.Close()

	_, err = conn.Do("EXISTS", "name")
	if err != nil {
		fmt.Println("redis exist failed", err)
	}

	nameStr, _ := redis.Bool(conn.Do("EXISTS", "name")) //int64转换成bool值
	fmt.Println(nameStr)
}
