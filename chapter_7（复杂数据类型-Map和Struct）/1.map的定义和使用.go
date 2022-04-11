package main

import (
	"fmt"
)

func main() {
	//map定义格式 map[keyType]valueType
	var dict1 map[int]string

	dict2 := make(map[int]string, 3)

	dict3 := map[[2]int][3]string{{0, 0}: {"a", "b", "c"}, {0, 1}: {"d", "e", "f"}, {0, 2}: {"g", "h", "i"}}

	fmt.Println(dict1, len(dict1)) //可以用len()查看键值对的个数，但是无法使用cap()查看容量
	fmt.Println(dict2, len(dict2))
	fmt.Println(dict3, len(dict3))
}
