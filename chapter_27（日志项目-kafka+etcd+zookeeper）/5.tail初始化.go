package main

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

func main() {
	var configObj = new(Config)
	//2.根据配置中的日志路径使用tail收集日志
	err := InitTail(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logrus.Error("init tailfile failed, err : %v", err)
		return
	}
}

var tailObj *tail.Tail

func InitTail(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 //日志自动切割 重新打开
		Follow:    true,                                 //日志自动切割 跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的末尾（os.SEEK）读数据
		MustExist: false,                                //文件不存在就报错
		Poll:      true,                                 //轮询
	}
	//打开文件开始读取数据
	tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		logrus.Error("tailfile: create tailObj for path:%s failed, err:%s\n", filename, err)
		return
	}
	return
}
