package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//基于sarama第三方库开发kafka client

func main() {
	//配置命令
	config := sarama.NewConfig()
	//发送完数据需要leader和follwer都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true

	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("This is a test log")

	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	fmt.Println("连接kafka成功")
	defer func(client sarama.SyncProducer) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed ,err:", err)
		return
	}
	fmt.Printf("pid:%v  offset:%v\n", pid, offset)
}
