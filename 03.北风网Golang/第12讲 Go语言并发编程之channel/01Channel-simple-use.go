package main

import (
	"fmt"
	"strconv"
	"time"
)

func Add(x, y int, quit chan int) {
	z := x + y
	fmt.Println(z)
	quit <- 1
}

func Read(ch chan int) {
	fmt.Println("读协程运行")
	value := <-ch
	fmt.Println("value:" + strconv.Itoa(value))
}

func Write(ch chan int) {
	fmt.Println("写协程运行")
	time.Sleep(4*time.Second)
	ch <- 10
}

func main() {
	ch := make(chan int)
	go Read(ch)
	go Write(ch)
	time.Sleep(10*time.Second)
	fmt.Println("主程序结束")
}
