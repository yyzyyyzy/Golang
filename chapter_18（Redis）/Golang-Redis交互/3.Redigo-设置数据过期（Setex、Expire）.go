package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
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

	_, err = conn.Do("SET", "age", "18")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = conn.Do("SETEX", "gender", "8", "男") //setex key 过期时间 value
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = conn.Do("EXPIRE", "age", "8") // expire key 过期时间
	if err != nil {
		fmt.Println("redis expire failed:", err)
	}

	reply1, err := redis.String(conn.Do("GET", "gender"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get gender: %v \n", reply1)
	}

	time.Sleep(8 * time.Second)

	reply2, err := redis.String(conn.Do("GET", "gender"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get gender: %v \n", reply2)
	}
}
