package main

import (
	"flag"
	"fmt"
	"strings"
)

var mac string

func main() {
	flag.StringVar(&mac, "mac", "xx:xx:xx:xx", "please enter your mac")
	flag.Parse()

	//将mac地址转换为不包含 ":" 和 "-" 的字符串
	strMac := ConvertChar(mac)
	if len(strMac) > 12 {
		panic("The correct format of the MAC address is: xx:xx:xx:xx")
	}
	macSlice := []byte(strMac)

	slice_1 := string(macSlice[0:2])
	slice_2 := string(macSlice[2:4])
	slice_3 := string(macSlice[4:6])
	slice_4 := string(macSlice[6:8])
	slice_5 := string(macSlice[8:10])
	slice_6 := string(macSlice[10:12])

	print(fmt.Sprintf(slice_1 + ":" + slice_2 + ":" + slice_3 + ":" + slice_4 + ":" + slice_5 + ":" + slice_6))
}

func ConvertChar(mac string) string {

	mac = strings.ToLower(mac)

	if strings.Contains(mac, "-") {
		mac = strings.ReplaceAll(mac, "-", "")

	}
	if strings.Contains(mac, ":") {
		mac = strings.ReplaceAll(mac, ":", "")
	}
	return mac
}
