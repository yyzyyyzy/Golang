package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsondataMap := `{"Age":18,"Name":"LZK","Sex":"男","hobby":["钢琴","摄影","潮流"]}`
	dataMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsondataMap), &dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dataMap)
}
