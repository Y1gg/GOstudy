package main

import (
	"fmt"
	"github.com/Y1gg/day80/logagent/conf"
	"github.com/Y1gg/day80/logagent/etcd"
	"github.com/Y1gg/day80/logagent/kafka"
	"gopkg.in/ini.v1"
	"time"
)

/*func run() {
	//1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)

		}
	}
}*/

// log-agent 入口程序
var (
	cfg = new(conf.AppConf)
)

func main() {
	//0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed ,err:%v\n", err)
		return
	}
	//cfg, err := ini.Load("./conf/config.ini")
	//1.初始化kafka的连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("初始化kafka连接成功")

	//2. 初始化ETCD
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("初始化etcd连接成功")

	//2.1、从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	//2.2、派一个哨兵去监视日志收集项的变化（有变化及时通知我的logagent实现热加载配置 ）
	if err != nil {
		fmt.Printf("etcd.GetConf failed ,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success.%v\n", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("index:%v\tvalue%v\n", index, value)
	}

	/*//2.打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("init taillog failed ,err:%v\n", err)
		return
	}
	fmt.Println("打开日志文件成功")
	run()*/

}
