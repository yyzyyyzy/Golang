package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	err = conn.Send("auth", "root")
	if err != nil {
		fmt.Println("conn.Send err", err)
	}
	defer conn.Close()

	_, err = conn.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = conn.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = conn.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	values, _ := redis.Values(conn.Do("lrange", "runoobkey", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}
