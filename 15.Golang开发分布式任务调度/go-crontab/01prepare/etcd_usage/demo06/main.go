package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	config         clientv3.Config
	client         *clientv3.Client
	err            error
	lease          clientv3.Lease
	leaseGrantResp *clientv3.LeaseGrantResponse
	leaseId        clientv3.LeaseID
	putResp        *clientv3.PutResponse
	getResp        *clientv3.GetResponse
	keepResp       *clientv3.LeaseKeepAliveResponse
	keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
	kv             clientv3.KV
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

	// 申请一个租约
	lease = clientv3.NewLease(client)

	// 申请一个10s的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	// 获取续约的ID
	leaseId = leaseGrantResp.ID

	// 5s后自动取消自动续租
	//ctx,_ :=context.WithTimeout(context.Background(),5*time.Second)
	//if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
	// 普通续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约答应的协程
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效")
					goto EXIT
				} else {
					fmt.Println("收到自动续租的应答:", keepResp.ID)
				}
			}
		}
	EXIT:
	}()

	kv = clientv3.NewKV(client)

	// Put一个KV, 让它与租约关联起来, 从而实现10秒后自动过期
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时检查key过期了没有
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("KV过期了")
			break
		}
		fmt.Println("还没有过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
