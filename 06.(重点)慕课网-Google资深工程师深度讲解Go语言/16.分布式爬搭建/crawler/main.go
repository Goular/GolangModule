package main

import (
	"baseLearn/16.分布式爬搭建/crawler/config"
	"baseLearn/16.分布式爬搭建/crawler/engine"
	"baseLearn/16.分布式爬搭建/crawler/persist"
	"baseLearn/16.分布式爬搭建/crawler/scheduler"
	"baseLearn/16.分布式爬搭建/crawler/zhenai/parser"
)

// 单机版
func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
