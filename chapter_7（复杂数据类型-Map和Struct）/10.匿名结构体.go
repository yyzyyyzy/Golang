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

	jsonData := getJsonData()
	Touch := struct {
		isTouch bool
	}{}
	json.Unmarshal(jsonData, &Touch)
	fmt.Println("解析（分离）部分结构:", Touch)

}
