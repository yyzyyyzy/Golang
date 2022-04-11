package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonDecodeMap() {
	path := "E:\\golandlearning\\chapter_8（复杂数据类型-json）\\map编码json.json"
	file, _ := os.Open(path)
	defer file.Close()

	dataMap := make(map[string]interface{})

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&dataMap)
	if err != nil {
		fmt.Println("jsonMap解码失败")
		return
	} else {
		fmt.Println("jsonMap解码成功")
		fmt.Println(dataMap)
	}
}

func JsonDecodeSlice() {
	path := "E:\\golandlearning\\chapter_8（复杂数据类型-json）\\slice编码json.json"
	file, _ := os.Open(path)
	defer file.Close()

	dataSlice := make([]interface{}, 0, 200)

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&dataSlice)
	if err != nil {
		fmt.Println("jsonSlice解码失败")
	} else {
		fmt.Println("jsonSlice解码成功")
		fmt.Println(dataSlice)
	}
}

func main() {
	JsonDecodeMap()
	JsonDecodeSlice()
}
