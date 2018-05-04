package main

import (
	"log"

	"fmt"

	"flag"

	"baseLearn/16.分布式爬搭建/crawler/config"
	"baseLearn/16.分布式爬搭建/distributed/persist"
	"baseLearn/16.分布式爬搭建/distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

// 创建启动配置参数,第三个参数就是我们的解析参数解析语句,使用go run xxxx.go --help即可查询
var port = flag.Int("port", 0, "The port for me to Listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must Specify a Port")
		return
	}
	//todo: 一般来说serveRpc方法是不会出错的，出错肯定是大错(系统出错)，所以可以直接错误退出就可以了

	// 常用的错误写法
	//err := serveRpc(":1234", "dating_profile")
	//if err != nil {
	//	panic(err)
	//}

	// 偷懒的错误写法
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
