package main

import (
	"encoding/json"
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

	key := "profile"
	imap := map[string]string{"Age": "18", "Name": "LZK", "Sex": "男"}
	value, _ := json.Marshal(imap)
	n, err := conn.Do("setnx", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["Age"])
	fmt.Println(imapGet["Sex"])

}
