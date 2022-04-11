package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main() {

	wait.Add(num) //等待一个协程
	for i := 0; i < 5; i++ {
		go ReadMyFile(&wait)
	}
	wait.Wait()
}

var wait sync.WaitGroup

const num = 5

func ReadMyFile(wait *sync.WaitGroup) {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\7.method是什么？方法表达式、方法变量==FunctionValue？"
	file, err := os.Open(path)
	HandleErr(err, "os.Open")
	defer file.Close()

	filebytes, err := ioutil.ReadAll(file)
	HandleErr(err, "ReadAll")
	fmt.Println(string(filebytes))
	wait.Done()
}

func HandleErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
