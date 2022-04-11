package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type S struct {
	Id     int
	Name   string
	Gender string
}

type C struct {
	Title   string
	Student []S
}

func StructEncodeJson() {
	class := C{Title: "清华班", Student: make([]S, 0, 200)}

	for i := 0; i < 20; i++ {
		stu := S{
			Id:     i,
			Name:   fmt.Sprintf("Student%d", i),
			Gender: "男",
		}
		class.Student = append(class.Student, stu)
	}

	path := "E:\\golandlearning\\chapter_8（复杂数据类型-json）\\struct编码json.json"
	dstfile, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer dstfile.Close()

	encoder := json.NewEncoder(dstfile)
	err := encoder.Encode(class)
	if err != nil {
		fmt.Println("struct编码失败")
		return
	} else {
		fmt.Println("struct编码成功")
	}
}

func MapEncodeJson() {
	dataMap := make(map[string]interface{})
	dataMap["Name"] = "LZK"
	dataMap["Age"] = 18
	dataMap["Sex"] = "男"
	dataMap["hobby"] = []string{"钢琴", "摄影", "潮流"}

	path := "E:\\golandlearning\\chapter_8（复杂数据类型-json）\\map编码json.json"
	dstfile, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer dstfile.Close()

	encoder := json.NewEncoder(dstfile)
	err := encoder.Encode(dataMap)
	if err != nil {
		fmt.Println("map编码失败")
		return
	} else {
		fmt.Println("map编码成功")
	}

}

func SliceEncodeJson() {
	dataSlice := []interface{}{"LZK", 18, [3]int{110, 120, 119}}

	path := "E:\\golandlearning\\chapter_8（复杂数据类型-json）\\slice编码json.json"
	dstfile, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer dstfile.Close()

	encoder := json.NewEncoder(dstfile)
	err := encoder.Encode(dataSlice)
	if err != nil {
		fmt.Println("slice编码失败")
		return
	} else {
		fmt.Println("slice编码成功")
	}
}

func main() {
	StructEncodeJson()
	MapEncodeJson()
	SliceEncodeJson()
}
