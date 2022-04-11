package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleRedisError(err, "redis.Dial")

	err = conn.Send("auth", "root")
	HandleRedisError(err, "conn.send")

	defer conn.Close()
}

func HandleRedisError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
