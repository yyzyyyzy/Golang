package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("notepad", "E:\\golandlearning\\chapter_0（底层原理）\\2.闭包怎么说，闭包")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
