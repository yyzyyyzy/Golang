package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//注册回调函数
	http.HandleFunc("/shein/hello", handler1)
	http.HandleFunc("/shein/bye", handler2)
	http.HandleFunc("/shein", handler3)

	//绑定tcp监听地址，并开始接受请求，然后调用服务端处理程序来处理传入的连接请求。
	//第一个参数 addr 即监听地址；第二个参数表示服务端处理程序，通常为nil
	//当参2为nil时，服务端调用 http.DefaultServeMux（路由器） 进行处理
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// ResponseWriter是一个接口，给客户端回发数据，Request是一个结构体，用来接收客户端发送的数据
func handler1(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method = ", request.Method) // 请求方法
	fmt.Println("URL = ", request.URL)       // 浏览器发送请求文件路径
	fmt.Println("header = ", request.Header) // 请求头
	fmt.Println("body = ", request.Body)     // 请求包体
	fmt.Println(request.RemoteAddr, "连接成功")  // 客户端网络地址

	writer.Write([]byte("<h1>你好，帅气的你又来了</h1>"))
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(request.RemoteAddr + "连接成功\n"))
	writer.Write([]byte("哥哥再见"))
}

// 打开SHEIN首页
func handler3(writer http.ResponseWriter, request *http.Request) {
	bytes, _ := ioutil.ReadFile("E:\\golandlearning\\chapter_16（网络编程）\\5.网络编程-HTTP通信\\shein.html")
	writer.Write(bytes)
}
