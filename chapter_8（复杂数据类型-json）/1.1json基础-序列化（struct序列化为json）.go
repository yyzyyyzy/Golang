package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Gender string
}

type Class struct {
	Title   string
	Student []Student
}

func main() {

	// 创建班级
	class := Class{Title: "南通班", Student: make([]Student, 0, 200)}

	// 创建20个同学
	for i := 1; i <= 200; i++ {
		stu := Student{
			Id:     i,
			Name:   fmt.Sprintf("Student%d", i),
			Gender: "男",
		}
		class.Student = append(class.Student, stu)
	}

	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println("创建失败")
		return
	}
	fmt.Printf("%s", data)
}
