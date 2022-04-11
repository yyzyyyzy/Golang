package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println()
	exec.Command("notepad").Run()
}
