package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// ETCD client put/get demo
// use etcd/clientv3

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

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "yue", "beautiful")
	cancel()

	if err != nil {
		fmt.Printf("put to etcd failed,err:\n", err)
		return
	}

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "yue", clientv3.WithPrefix())
	resp, err := cli.Get(ctx, "yue")
	cancel()

	if err != nil {
		fmt.Printf("get from etcd failed,err:%v\n", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
