package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanWord() {
	str := "A vegan woman in Australia is taking her neighbor to court because he keeps -- wait for it -- barbecuing stuff in his backyard. The smell of the meat and fish he's been cooking has apparently prevented her from enjoying her own backyard."
	scanner := bufio.NewScanner(strings.NewReader(str))

	// 设置分词方式(按行读取)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println("-------------------------------------------------")

	// 按单词读取
	wordScanner := bufio.NewScanner(strings.NewReader(str))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		fmt.Println(wordScanner.Text())
	}
}

func ScanFile() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\14.1goroutine的执行过程"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open closed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	ScanWord()
	//ScanFile()
}
