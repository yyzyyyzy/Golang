package main

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	Id     int
	Name   string
	Gender string
}

type Class1 struct {
	Title   string
	Student []*Student1
}

func main() {
	jsondataStruct := `{"Title":"南通班","Student":[{"Id":1,"Name":"Student1","Gender":"男"},{"Id":2,"Name":"Student2","Gender":"男"},{"Id":3,"Name":"Student3","Gender":"男"},{"Id":4,"Name":"Student4","Gender":"男"},{"Id":5,"Name":"Student5","Gender":"男"},{"Id":6,"Name":"Student6","Gender":"男"},{"Id":7,"Name":"Student7","Gender":"男"},{"Id":8,"Name":"Student8","Gender":"男"},{"Id":9,"Name":"Student9","Gender":"男"},{"Id":10,"Name":"Student10","Gender":"男"},{"Id":11,"Name":"Student11","Gender":"男"},{"Id":12,"Name":"Student12","Gender":"男"},{"Id":13,"Name":"Student13","Gender":"男"},{"Id":14,"Name":"Student14","Gender":"男"},{"Id":15,"Name":"Student15","Gender":"男"},{"Id":16,"Name":"Student16","Gender":"男"}]}`
	dataStruct := &Class1{}
	err := json.Unmarshal([]byte(jsondataStruct), dataStruct)
	if err != nil {
		fmt.Println("创建失败")
		return
	}
	fmt.Println(dataStruct)
	fmt.Println(dataStruct.Student[15].Name)
}
