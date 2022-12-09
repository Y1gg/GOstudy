package main

//watch 操作
//watch用来获取未来更改的通知

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//watch

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		//handle error!
		fmt.Printf("connect to etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	defer cli.Close()

	//watch key : yue change
	//派一个哨兵一直监视着yue这个key的变化（新增、修改、删除）
	rch := cli.Watch(context.Background(), "yue") //<-chan WatchResponse
	//从通道尝试取值（监视的信息）
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%s\tKey:%s\tValue:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
