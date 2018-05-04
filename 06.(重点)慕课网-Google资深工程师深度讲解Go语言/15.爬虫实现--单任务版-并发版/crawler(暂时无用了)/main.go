package main

import (
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/engine"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/persist"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/scheduler"
	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/zhenai/parse"
)

func main() {
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parse.ParseCity,
	//})
}
