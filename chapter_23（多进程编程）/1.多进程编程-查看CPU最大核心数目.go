package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(runtime.NumCPU())
}
