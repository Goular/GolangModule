package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan int)
	//协程1
	go func() {
		for i := 1; i < 100; i++ {
			if i == 10 {
				//主动出让cpu 使用的话 需要 导入 runtime包
				//runtime.Gosched()

				//让出执行权2阻塞
				<-ch
			}
			fmt.Println("routine 1:" + strconv.Itoa(i))

		}
	}()

	//协程2
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("routine 2:" + strconv.Itoa(i))
		}
		ch <- 1
	}()

	time.Sleep(time.Second)
}
