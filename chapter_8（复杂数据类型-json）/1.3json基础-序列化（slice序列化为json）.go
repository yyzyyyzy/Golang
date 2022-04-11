package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dataSlice := []interface{}{"LZK", 18, [3]int{110, 120, 119}}
	bytes, err := json.Marshal(dataSlice)
	if err != nil {
		fmt.Println(bytes)
	}
	fmt.Println(string(bytes))
}
