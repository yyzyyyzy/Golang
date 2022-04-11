package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := `D:\kafka\kafkalog\web_log-0\00000000000000000000.log`
	config := tail.Config{
		ReOpen:    true,                                 //日志自动切割 重新打开
		Follow:    true,                                 //日志自动切割 跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的末尾（os.SEEK）读数据
		MustExist: false,                                //文件不存在就报错
		Poll:      true,                                 //轮询
	}
	//打开文件开始读取数据
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tail: %s is failed, err:%s\n", filename, err)
		return
	}

	//开始读取数据
	var (
		msg *tail.Line
		ok  bool
	)

	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second) //读取出错等待1s
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
