package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println就是os.Stdout.Write()的封装
	fmt.Println("你好，golang")
	os.Stdout.WriteString("你好，golang\n")
	os.Stdout.Write([]byte("你好，golang\n"))
}
