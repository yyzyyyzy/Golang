package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Id   int64  `json:"id,string"`
	Name string `json:"name"`
}

type Info struct {
	Token string `json:"token"`
}

func AnonymousStruct(user interface{}) {
	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("解析失败", err)
	}
	fmt.Println(string(bytes))
}

func main() {
	u1 := struct {
		UserInfo
		Info
		HasHash bool
	}{
		UserInfo: UserInfo{Id: 123, Name: "LZK"},
		Info:     Info{Token: "dsd34sf146541f65xc4a6s"},
		HasHash:  true,
	}
	AnonymousStruct(u1)
}
