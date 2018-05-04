package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	go func() {
		for n := range w.in {
			fmt.Printf("Worker %d received %c \n", id, n)
			w.done()
		}
	}()
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// 因为需要引用主协程的wg，所以必须是指针变量
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(20)
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for i, workder := range workers {
		workder.in <- 'a' + i
	}
	for i, workder := range workers {
		workder.in <- 'A' + i
	}
}

func main() {
	chanDemo()
}
