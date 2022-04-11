package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsondataSlice := `["LZK",18,[110,120,119]]`

	dataSlice := make([]interface{}, 0)

	err := json.Unmarshal([]byte(jsondataSlice), &dataSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dataSlice)
}
