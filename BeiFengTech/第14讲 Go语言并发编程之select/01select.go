package main

import (
	//"time"
	"fmt"
	"time"
)

/**
	channel的使用
 */
func main() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1
	}(ch)
	time.Sleep(time.Second)
	select {
	case <-ch:
		fmt.Println("Come to read ch！")
	default:
		fmt.Println("Come to default！")
	}
	fmt.Println("Come to Main End！")
}
