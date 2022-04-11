package main

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func main() {
	var configObj = new(Config)
	err := InitKafka([]string{configObj.KafkaConfig.Address})
	if err != nil {
		logrus.Error("init kafka failed, err : %v", err)
		return
	}
	logrus.Info("init kafka success!")
}

func InitKafka(address []string) (err error) {

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
