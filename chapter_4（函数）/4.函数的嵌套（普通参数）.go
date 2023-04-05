package main

import "fmt"

// 发送邮件
func send_msg() {
}

// 校验函数
func check_info(username, passwd, useremail string) {
}

// 注册函数
func regsiter() {
	check_info("LZK", "123456", "916990143@qq.com")
	send_msg()
	fmt.Println("用户注册成功")
}

func main() {
	regsiter()
}
