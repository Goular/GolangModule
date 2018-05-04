package main

import (
	"fmt"
)

/**
 *
 */
func doWorker(id int, c chan int, done chan bool) {
	go func() {
		for n := range c {
			fmt.Printf("Worker %d received %c \n", id, n)
			// 这种写法会产生死锁
			done <- true
			//go func() {
			//	done <- true
			//}()
		}
	}()
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, workder := range workers {
		workder.in <- 'a' + i
		// 不管收什么，只要可以收内容，说明已经可以完毕程序了
		// <-workers[i].done
	}
	for _, worker := range workers {
		<-worker.done
	}
	for i, workder := range workers {
		workder.in <- 'A' + i
		// 不管收什么，只要可以收内容，说明已经可以完毕程序了
		// <-workers[i].done
	}
	for _, worker := range workers {
		<-worker.done
	}
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}

	// 有了等待channel的内容，那么下面的sleep就可以关闭，一直让cpu hold住是不好的
	//time.Sleep(10 * time.Second)
}

func main() {
	chanDemo()
}
