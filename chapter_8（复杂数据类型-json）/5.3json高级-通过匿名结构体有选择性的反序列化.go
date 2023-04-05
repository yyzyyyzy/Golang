package main

import (
	"encoding/json"
	"fmt"
)

// 定义手机屏幕
type Screen01 struct {
	Size       float64 //屏幕尺寸
	ResX, ResY int     //屏幕分辨率 水平 垂直
}

// 定义电池容量
type Battery struct {
	Capacity string
}

// 返回json数据
func getJsonData() []byte {
	//tempData 接收匿名结构体（匿名结构体使得数据的结构更加灵活）
	tempData := struct {
		Screen01
		Battery
		HashTouchId bool // 是否有指纹识别
	}{
		Screen01:    Screen01{Size: 12, ResX: 36, ResY: 36},
		Battery:     Battery{"6000毫安"},
		HashTouchId: true,
	}
	jsonData, _ := json.Marshal(tempData) //将数据转换为json
	return jsonData
}
func main() {
	jsonData := getJsonData() //获取json数据
	fmt.Printf("%s\n", jsonData)
	fmt.Println("=========解析（分离）出的数据是===========")
	//自定义匿名结构体，解析（分离）全部数据
	allData := struct {
		Screen01
		Battery
		HashTouchId bool
	}{}
	json.Unmarshal(jsonData, &allData)
	fmt.Println("解析（分离）全部结构为：", allData)
	//自定义匿名结构体，通过json数据，解析（分离）对应的结构（可以是部分结构）
	screenBattery := struct {
		Screen01
		Battery
	}{}
	json.Unmarshal(jsonData, &screenBattery) //注意：此处只能为结构体指针（一般参数为interface{}，都采用地址引用（即地址传递））
	fmt.Println("解析（分离）部分结构:", screenBattery)
	//自定义匿名结构体，解析（分离）部分结构
	batteryTouch := struct {
		Battery
		isTouch bool
	}{}
	json.Unmarshal(jsonData, &batteryTouch)
	fmt.Println("解析（分离）部分结构:", batteryTouch)
	//自定义匿名结构体，解析（分离）部分不存在的结构
	temp1 := struct {
		Battery
		Detail struct {
			Name  string
			Price uint16
		}
	}{}
	json.Unmarshal(jsonData, &temp1)
	fmt.Println("解析（分离）部分不存在的结构", temp1)
	//自定义匿名结构体，解析（分离）完全不存在的结构
	temp2 := struct {
		User  string
		Price uint16
	}{}
	json.Unmarshal(jsonData, &temp2)
	fmt.Println("解析（分离）完全不存在的结构:", temp2)
}
