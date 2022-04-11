package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//客户端访问web服务器数据，主要使用func Get(url string) (resp *Response, err error)函数来完成。
//读到的响应报文数据被保存在 Response 结构体中。
func Get() {
	resp, err := http.Get("http://www.baidu.com/s?wd=药水哥")
	HandleHTTPClientErr(err, "http.Get")

	defer resp.Body.Close()

	//读取响应体
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleHTTPClientErr(err, "ioutil.ReadAll")
	fmt.Println("网页内容为：", string(bytes))
}

//服务器发送的响应包体被保存在Body中。可以使用它提供的Read方法来获取数据内容。保存至切片缓冲区中，拼接成一个完整的字符串来查看。
//结束的时候，需要调用Body中的Close()方法关闭io。
func Post() {
	url := "http://httpbin.org/post?name=LZK&age=18"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("tall=187&weight=150"))
	HandleHTTPClientErr(err, "http.Post")

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HandleHTTPClientErr(err, "ioutil.ReadAll")
	fmt.Println("网页内容为：", string(bytes))
}

func main() {
	//Get()
	Post()
}

func HandleHTTPClientErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
