package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	// 要遍历的文件夹
	dir := "E:\\golandlearning"

	// 遍历的文件夹
	// 参数：要遍历的文件夹，层级（默认：0）
	findDir(dir, 0)

}

// 遍历的文件夹
func findDir(dir string, num int) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {

		// 重复输出制表符，模拟层级结构
		print(strings.Repeat("\t", num))

		// 判断是不是目录
		if fi.IsDir() {
			println(`目录：`, fi.Name())
			findDir(dir+`/`+fi.Name(), num+1)
		} else {
			println(`文件：`, fi.Name())
		}
	}
}
