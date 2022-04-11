package main

import "os"

func main() {
	os.Mkdir("1", 0666)                //文件夹创建
	os.MkdirAll("1/2/3/4/5/6/7", 0666) //多级文件夹创建
}
