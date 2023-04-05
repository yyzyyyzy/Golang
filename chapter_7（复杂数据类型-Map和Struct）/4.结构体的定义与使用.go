package main

import "fmt"

// 结构体定义在值函数外部
type good_people struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  int
}

func main() {
	//初始化结构体
	var human1 = good_people{
		Id:     1,
		Name:   "王浩川",
		Gender: "男",
		Age:    18,
		Phone:  13804056806,
	}

	//自动推导
	human2 := good_people{
		Id:     2,
		Name:   "李子康",
		Gender: "男",
		Age:    18,
		Phone:  13610850940,
	}

	fmt.Println(human1, human2)

	//结构体取值
	fmt.Println(human1.Name, human1.Phone)

	//结构体的赋值
	human1.Phone = 110
}
