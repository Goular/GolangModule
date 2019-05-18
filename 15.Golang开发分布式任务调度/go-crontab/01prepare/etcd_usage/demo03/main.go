package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	getResp *clientv3.GetResponse
)

func main() {
	// 构建etcd client的config
	config = clientv3.Config{
		Endpoints:   []string{"106.12.118.76:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 构建etcd client
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
	}
	kv = clientv3.NewKV(client)

	//if getResp, err = kv.Get(context.TODO(), "/crontab/jobs/job1", /*clientv3.WithCountOnly()*/); err != nil {
	if getResp, err = kv.Get(context.TODO(), "/crontab/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs, getResp.Count)
		for _, value := range getResp.Kvs {
			fmt.Println(value)
		}
	}
}
