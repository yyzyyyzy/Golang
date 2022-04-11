package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

//日志收集的客户端
//类似的有filebeat
//收集指定目录的日志文件，发送到kafka
func main() {
	//0.读取配置文件
	var configObj = new(Config)
	err := ini.MapTo(configObj, "E:\\golandlearning\\chapter_27（日志项目-kafka+etcd+zookeeper）\\config.ini")
	if err != nil {
		logrus.Error("load config failed, err : %v", err)
		return
	}
	fmt.Printf("%#v\n", configObj)

	//1.初始化kafka
	//2.根据配置中的日志路径使用tail收集日志
	//3.把日志发往kafka
}

type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
}
type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}
