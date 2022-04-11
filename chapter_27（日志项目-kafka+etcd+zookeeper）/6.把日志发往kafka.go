package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	err := run()
	if err != nil {
		logrus.Error("run failed, err : %v", err)
		return
	}
}

func run() (err error) {
	for {
		msg, ok := <-tailObj.Lines
		if !ok {
			logrus.Warn("tail file close reopen, filename:%s\n", tailObj.Filename)
			time.Sleep(time.Second) //读取出错等待1s
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
