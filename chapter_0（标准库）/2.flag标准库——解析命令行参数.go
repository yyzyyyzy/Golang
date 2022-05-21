package main

import (
	"flag"
	"fmt"
	"time"
)

func main1() {
	name := flag.String("userName", "lzk", "test for flag.string")
	age := flag.Int64("userAge", 18, "test for flag.Int")
	married := flag.Bool("isMarried", false, "test for flag.Bool")
	delay := flag.Duration("delay", 10*time.Second, "test for flag.Duration")

	flag.Parse()

	fmt.Println(*name, *age, *married, *delay)

	// go run .\2.flag标准库——解析命令行参数.go -userName whc -userAge 29 -isMarried true -delay=0.5ms
}

func main() {
	var name string
	var age int64
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "userName", "lzk", "test for flag.string")
	flag.Int64Var(&age, "userAge", 18, "test for flag.Int")
	flag.BoolVar(&married, "isMarried", false, "test for flag.Bool")
	flag.DurationVar(&delay, "delay", 10*time.Second, "test for flag.Duration")

	flag.Parse()

	fmt.Println(name, age, married, delay)

	// go run .\2.flag标准库——解析命令行参数.go -userName whc -userAge 29 -isMarried true -delay=0.5ms
}
