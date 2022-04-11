package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type MyData struct {
	ID     int64  `json:"-"`             //"-" 表示不进行序列化
	Number int64  `json:"number,string"` //,string表示为了保存精度改变序列化输出的类型
	Name   string `json:"name"`
	Phone  string `json:"phone,omitempty"` //omitempty表示如果实例中phone为零值或空值，那么不进行序列化
}

func main() {
	data1 := MyData{
		ID:     math.MaxInt64,
		Number: 18,
		Name:   "李子康",
		Phone:  "",
	}

	bytes, _ := json.Marshal(data1)
	fmt.Println(string(bytes)) //ID没被序列化，number输出类型改为了string，Phone为空值没被序列化

}
