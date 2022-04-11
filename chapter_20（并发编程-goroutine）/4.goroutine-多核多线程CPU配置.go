package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
)

var wait1 sync.WaitGroup

func main() {

	wait1.Add(1)

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu) //使用多核多线程优化

	go ReadMyFile1(&wait1)
	wait1.Wait() //等待
}

func ReadMyFile1(wait1 *sync.WaitGroup) {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\7.method是什么？方法表达式、方法变量==FunctionValue？"
	file, err := os.Open(path)
	HandleErr1(err, "os.Open")
	defer file.Close()

	filebytes, err := ioutil.ReadAll(file)
	HandleErr1(err, "ReadAll")
	fmt.Println(string(filebytes))
	wait1.Done()
}

func HandleErr1(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
