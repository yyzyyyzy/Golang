package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dataMap := make(map[string]interface{})
	dataMap["Name"] = "LZK"
	dataMap["Age"] = 18
	dataMap["Sex"] = "男"
	dataMap["hobby"] = []string{"钢琴", "摄影", "潮流"}

	bytes, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
