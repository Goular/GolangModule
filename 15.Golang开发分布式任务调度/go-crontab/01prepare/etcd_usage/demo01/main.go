package demo01

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// demo01 etcd golang 客户端连接
// 106.12.118.76
func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"106.12.118.76:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	client = client
	fmt.Println(client)
}
