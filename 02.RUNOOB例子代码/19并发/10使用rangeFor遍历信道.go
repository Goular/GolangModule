package main

import (
	"fmt"
)

/*
for range 不断从信道 ch 中接收数据，直到该信道关闭。一旦 ch 关闭，循环自动退出。
 */
func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ",v)
	}
}