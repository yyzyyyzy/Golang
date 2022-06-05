package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

//WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址
//WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量
//Add(n) 把计数器设置为n
//Done() 每次把计数器-1
//wait() 会阻塞代码的运行，直到计数器地值减为0

func main() {

	wait.Add(num)
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
