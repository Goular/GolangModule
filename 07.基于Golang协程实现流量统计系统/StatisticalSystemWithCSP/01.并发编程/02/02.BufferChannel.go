package main

import (
	"fmt"
	"strconv"
)

func sample(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "Hello Imooc！Num:" + strconv.Itoa(i)
	}
}

func sample2(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "Hello Sample 2！Num:" + strconv.Itoa(i)
	}
}

func main() {
	ch := make(chan string, 3)
	fmt.Println("---begin------")
	for i := 0; i < 3; i++ {
		go sample(ch)
		go sample2(ch)
	}

	// 普通写法
	//for i := 0; i < 18; i++ {
	//	fmt.Println(<-ch)
	//}

	//range写法,由于读多于写，会产生死锁
	for str := range ch {
		fmt.Println(str)
	}

	fmt.Println("--- end ------")
}
