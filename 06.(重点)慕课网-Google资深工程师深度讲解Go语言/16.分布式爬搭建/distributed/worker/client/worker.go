package client

import (
	"net/rpc"

	"baseLearn/16.分布式爬搭建/crawler/engine"
	cconfig "baseLearn/16.分布式爬搭建/distributed/config"
	"baseLearn/16.分布式爬搭建/distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}
	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(cconfig.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
