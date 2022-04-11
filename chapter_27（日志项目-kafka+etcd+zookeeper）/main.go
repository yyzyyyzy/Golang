package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"time"
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
	err = Init1([]string{configObj.KafkaConfig.Address})
	if err != nil {
		logrus.Error("init kafka failed, err : %v", err)
		return
	}
	logrus.Info("init kafka success!")
	//2.根据配置中的日志路径使用tail收集日志
	err = Init2(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logrus.Error("init tailfile failed, err : %v", err)
		return
	}
	logrus.Info("init tailfile success!")
	//3.把日志发往kafka
	err = run()
	if err != nil {
		logrus.Error("run failed, err : %v", err)
		return
	}
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

var Client sarama.SyncProducer

func Init1(address []string) (err error) {

	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	Client, err = sarama.NewSyncProducer(address, config) //kafka默认端口9092
	if err != nil {
		logrus.Error("kafka: producer closed, err:", err)
		return
	}
	return
}

var tailObj *tail.Tail

func Init2(filename string) (err error) {
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
