package main

import "fmt"

func main() {
	dict1 := map[string]string{"name": "lzk", "age": "18", "phone": "13951801007", "email": "916990143@qq.com"}

	//增删改查
	//1.增/改
	dict1["email"] = "857644244@qq.com"

	//2.删
	delete(dict1, "phone")

	//3.查
	for k, v := range dict1 {
		fmt.Println(k, v)
	}
}
