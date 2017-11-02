package main

import (
	"fmt"
	"time"
)

/**
	使用time.After来替代timeout channel的使用
 */
func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("Come to read ch！")
	case <-time.After(time.Second):
		fmt.Println("Come to time out！")
	}
	time.Sleep(time.Second)
	fmt.Println("Come to Main End！")
}
