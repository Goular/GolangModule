package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	delResp *clientv3.DeleteResponse
	kvpair  *mvccpb.KeyValue
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
	// 删除指定的kv
	if delResp, err = kv.Delete(context.TODO(), "/crontab/jobs/job1", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(delResp.PrevKvs)
	fmt.Println(delResp.Deleted)
	// 被删除之前的value是什么
	if len(delResp.PrevKvs) != 0 {
		for _, kvpair = range delResp.PrevKvs {
			fmt.Println("删除了:", string(kvpair.Key), string(kvpair.Value))
		}
	}
}
