package main

import (
	"encoding/xml"
	"fmt"
	"net"
	"time"
)

type JiaoHangSms struct {
	Header Header
	Body   Body
}

// 定义报文头
type Header struct {
	PacLen    int    // 报文长度
	ChannNo   string // 请求渠道
	ReqTime   string // 请求时间
	TransCode string // 交易码
}

// 定义报文体
type Body struct {
	TemNo     string // 模板编号
	SendFlag  string // 预约标志
	SendTime  string // 预约时间
	MsgType   string // 消息标志类型
	ActNo     string // 卡号/账号
	CustNo    string // 客户号
	OrgNo     string // 机构号
	ObjAddr   string // 目标地址
	ParamList ParamList
}

type ParamList struct {
	Params []Param `xml:"Param"`
}

type Param struct {
	ArgName  string `xml:"ArgName"`
	ArgValue string `xml:"ArgValue"`
}

// 定义消息
type Message struct {
	XMLName xml.Name `xml:"Message"`
	Body    Body
}

func main() {
	// 定义header
	header := &Header{
		PacLen:    0,
		ChannNo:   "CHNGEMS",
		ReqTime:   time.Now().Format("20060102150405"),
		TransCode: "MSCSMS0016",
	}

	// 定义body
	body := &Body{
		TemNo:    "124027",
		SendFlag: "1",
		SendTime: "",
		MsgType:  "01",
		ActNo:    "6222520114933550",
		CustNo:   "100001",
		OrgNo:    "0000000",
		ObjAddr:  "9502230",
		ParamList: ParamList{
			Params: []Param{
				{
					ArgName:  "securityCode",
					ArgValue: "123456",
				},
			},
		},
	}

	// 将body序列化为xml字符串
	bodyXml, err := xml.Marshal(body)
	if err != nil {
		panic(err)
	}

	// 计算报文长度
	header.PacLen = len(bodyXml) + 8 + 7 + 14 + 20

	// 将header和body组成报文
	message := make([]byte, header.PacLen)
	copy(message[0:8], []byte(fmt.Sprintf("%08d", header.PacLen)))
	copy(message[8:15], []byte(header.ChannNo))
	copy(message[15:29], []byte(header.ReqTime))
	copy(message[29:49], []byte(header.TransCode))
	copy(message[49:], bodyXml)

	// 发送报文
	conn, err := net.Dial("tcp", "localhost:8080") // 替换成实际的地址和端口
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sent message:", string(message))
}
