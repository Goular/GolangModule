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
	putResp *clientv3.PutResponse
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
	if putResp, err = kv.Put(context.TODO(), "/crontab/jobs/job3", "byebye", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("PrevValue:", string(putResp.PrevKv.Value))
		}
	}
}
