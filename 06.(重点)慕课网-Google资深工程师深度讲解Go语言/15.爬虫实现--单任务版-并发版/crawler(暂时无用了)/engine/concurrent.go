package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (engine *ConcurrentEngine) Run(seeds ...Request) {
	// 输入
	//in := make(chan Request)
	// 输出
	out := make(chan ParseResult)
	// 将chan送入调度器
	//engine.Scheduler.ConfigureMasterWorkerChan(in)

	engine.Scheduler.Run()

	// 创建ITEM
	for i := 0; i < engine.WorkerCount; i++ {
		// createWorker(in, out)
		createWorker(engine.Scheduler.WorkerChan(), out, engine.Scheduler)
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request : %s", r.Url)
			continue
		}
		engine.Scheduler.Submit(r)
	}
	//profileCount := 0
	//itemCount := 0
	// 接收 out 出去的 item
	for {
		result := <-out
		for _, item := range result.Items {
			// item.(model.Profile)一般不要这样写，因为这种断言是放在外面的，而不是放在engine的
			// if _, ok := item.(model.Profile); ok {
			//log.Printf("Got item #%d : %v \n", itemCount, item)
			//itemCount++
			//}

			// todo:尽快将保存item的任务脱手
			go func() { engine.ItemChan <- item }()
		}
		// 这里进行URL去重操作
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request : %s", request.Url)
				continue
			}
			engine.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			log.Printf("返回结果:%v\n", result)
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
