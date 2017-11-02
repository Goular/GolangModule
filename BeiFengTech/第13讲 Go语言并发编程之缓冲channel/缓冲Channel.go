package main

import (
	"fmt"
	"time"
)

var ch chan int

func test_channel() {
	ch <- 1
	fmt.Println("ch 1")
	ch <- 1
	fmt.Println("ch 2")
	ch <- 1
	fmt.Println("come to end goroutine 1")
}

func main() {
	//默认第二个参数就是缓冲区的容量，默认不填为0
	//ch = make(chan int, 0) // 等价于 ch = make(chan int) 都是不带缓冲的channel
	//ch = make(chan int, 1) // 等价于 ch = make(chan int) 都是不带缓冲的channel
	ch = make(chan int, 2) // 是带缓冲的channel
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Println("running end!")
	<-ch

	time.Sleep(time.Second)
}
