package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("getCurrentAbPathByExecutable = ", getCurrentAbPathByExecutable())
}

// 获取当前执行程序所在的绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}
