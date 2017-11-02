package main

import (
	//"time"
	"fmt"
	"time"
)

/**
	channel的使用
	由于select仅能做一次，所以ch的channel不执行，因超时走大了timeout，然后结束整个流程
 */
func main() {
	ch := make(chan int)
	timeout := make(chan int)
	go func(ch chan int) {
		ch<- 1
	}(timeout)
	go func(ch chan int) {
		time.Sleep(2 * time.Second)
		ch <- 1
	}(ch)
	time.Sleep(time.Second)
	select {
	case <-ch:
		fmt.Println("Come to read ch！")
	case <-timeout:
		fmt.Println("Come to time out！")
	default:
		fmt.Println("Come to default！")
	}
	fmt.Println("Come to Main End！")
}
