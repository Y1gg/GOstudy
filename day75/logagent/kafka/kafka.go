package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//专门往kafka写日志的模块

var (
	client sarama.SyncProducer
	//声明一个全局的连接kafka的生产者 client
)

// Init 初始化client
func Init(addrs []string) (err error) {
	//配置命令
	config := sarama.NewConfig()
	//发送完数据需要leader和follwer都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	return
}

func SendToKafka(topic, data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	//发送到kafka
	pid, offset, err := client.SendMessage(msg)
	fmt.Println("xxx")
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
