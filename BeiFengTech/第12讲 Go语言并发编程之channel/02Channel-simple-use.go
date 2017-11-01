package main

import (
	"fmt"
	"strconv"
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
	ch <- 10
}

func main() {
	chs := make([]chan int, 10)
	length := len(chs)


	fmt.Println("length:",length)

	for i := 0; i <length; i++ {
		chs[i] = make(chan int)
		go Add(i,i,chs[i])
	}
	for _,ch := range chs{
		<-ch
	}
}
