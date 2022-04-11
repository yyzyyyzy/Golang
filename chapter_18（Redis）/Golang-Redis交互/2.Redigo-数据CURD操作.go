package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleRedisErr(err, "redis.Dial")
	err = conn.Send("auth", "root")
	HandleRedisErr(err, "conn.send")
	defer conn.Close()

	//增、改
	_, err = conn.Do("SET", "name", "bill")
	HandleRedisErr(err, "conn.Do set")
	//查
	_, err = conn.Do("GET", "name")
	HandleRedisErr(err, "conn.Do get")
	//删
	_, err = conn.Do("DEL", "name")
	HandleRedisErr(err, "conn Do del")

	nameStr, _ := redis.String(conn.Do("GET", "name")) //[]uint8需要转换为string类型
	fmt.Println(nameStr)
}

func HandleRedisErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
