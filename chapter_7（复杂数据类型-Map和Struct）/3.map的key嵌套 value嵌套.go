package main

import "fmt"

func main() {
	//key 嵌套 键不重复/可哈希（整形/浮点型/布尔型/数组）
	dict1 := map[[2]int]string{{0, 1}: "lzk", {0, 2}: "whc"}
	fmt.Println(dict1)

	//value 嵌套
	dict2 := make(map[string]map[int]int)
	dict2["age1"] = map[int]int{0: 99, 100: 199}
	dict2["age2"] = map[int]int{0: 99, 100: 199}
	fmt.Println(dict2)

	dict3 := make(map[string][2]map[string]string) //数组中嵌套map（map的值为数组，数组的元素确定）
	dict3["name"] = [2]map[string]string{map[string]string{"name": "lzk", "age": "18"}, map[string]string{"name": "whc", "age": "28"}}
	fmt.Println(dict3)

	dict4 := make(map[string][]map[string]string) //切片中嵌套map（map的值为切片，切片的元素不确定）
	dict4["name"] = []map[string]string{map[string]string{"name": "lzk"}, map[string]string{"name": "whc", "age": "18"}}
	fmt.Println(dict4)
}
