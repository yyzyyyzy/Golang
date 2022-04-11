package main

import "fmt"

func main() {
	//数组初始化
	var arr1 [10]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [5]int{2: 10, 4: 20} //下标为2，值为10，下标为4，值为20

	//取数组长度
	arrlength1 := len(arr1)
	arrlength2 := len(arr2)
	arrlength3 := len(arr3)

	//打印
	fmt.Printf("数组为%v, 长度为%v\n", arr1, arrlength1)
	fmt.Printf("数组为%v, 长度为%v\n", arr2, arrlength2)
	fmt.Printf("数组为%v, 长度为%v\n", arr3, arrlength3)

	//取数组存储地址
	for i := 0; i < 5; i++ { //数组存储地址是连续的
		fmt.Println(&arr2[i])
	}
}
