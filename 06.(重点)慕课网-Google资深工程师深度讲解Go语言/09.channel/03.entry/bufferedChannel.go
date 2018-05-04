package main

import (
	"fmt"
	"time"
)

//func worker(id int, c <-chan int) {
//	for {
//		val, ok := <-c
//		if !ok {
//			break
//		}
//		fmt.Printf("Worker %d received %d \n", id, val)
//	}
//}

func worker(id int, c <-chan int) {
	for {
		for val := range c {
			fmt.Printf("Worker %d received %d \n", id, val)
		}
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelClose() {

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	close(c)
	time.Sleep(time.Second)
}

func main() {
	bufferedChannel()
}
