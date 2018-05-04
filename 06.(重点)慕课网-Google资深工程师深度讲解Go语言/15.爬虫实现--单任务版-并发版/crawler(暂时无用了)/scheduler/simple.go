package scheduler

import "baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- request }()
}
