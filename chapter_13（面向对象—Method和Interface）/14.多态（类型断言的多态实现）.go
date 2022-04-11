package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机正在打电话")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (pc Computer) PcWorking(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok { //struct也是类型，可以进行类型断言
		phone.Call()
	}
	usb.Stop()
}

func main() {
	usbArr := [2]Usb{Phone{name: "小米"}, Camera{name: "尼康"}}

	computer := Computer{}
	for _, v := range usbArr {
		computer.PcWorking(v)
	}
}
