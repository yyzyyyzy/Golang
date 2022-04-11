package main

import "fmt"

type friends struct {
	id   int
	name string
	age  int
}

//数组类型的结构体（值传递）
func bubbleSort1(arr [3]friends) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].age < arr[j+1].age {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

//切片类型的结构体（引用传递）
func bubbleSort2(arr []friends) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].age < arr[j+1].age {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	//数组类型的结构体（值传递）
	class1 := [3]friends{
		friends{1, "lzk", 18},
		friends{2, "whc", 19},
		friends{3, "dce", 20}}
	fmt.Println(class1)
	bubbleSort1(class1)

	//切片类型的结构体（引用传递）
	class2 := []friends{
		friends{4, "lzx", 21},
		friends{5, "lyj", 22},
		friends{6, "gqb", 23}}
	fmt.Println(class2)
	bubbleSort2(class2)
}
