package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			// 但是这样写还是存在阻塞的问题
			// w <- n
			// hasValue = true
			values = append(values, n)
		case n := <-c2:
			// 但是这样写还是存在阻塞的问题
			//w <- n
			// hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue: // 只要通道不阻塞，都是可以执行的
			// 这样写虽然难以理解，但是其实他们是平级的，所以上面如果都不执行，就会执行当前命令
			// hasValue = false
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Time out")
		case <-tick:
			fmt.Println("Queue Len = ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
