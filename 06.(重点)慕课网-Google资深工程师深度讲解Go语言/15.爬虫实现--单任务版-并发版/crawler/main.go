package main

import (
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler/config"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler/engine"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler/persist"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler/scheduler"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler/zhenai/parser"
)

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
