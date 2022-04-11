package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("E:\\webstorm\\WebStorm 2021.3.2\\bin\\webstorm64.exe")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
