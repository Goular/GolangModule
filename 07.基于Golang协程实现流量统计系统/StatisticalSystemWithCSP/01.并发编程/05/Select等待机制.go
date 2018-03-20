package main

import (
	"strconv"
	"time"
	"fmt"
)

func sample(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "I'm sample 1 num :" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
}

func sample2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)
	for i := 0; i < 10; i++ {
		go sample(ch1)
		go sample2(ch2)
	}
	for {
		select {
		case _str, ok := <-ch1:
			if !ok {
				fmt.Println("CH1 Failed!")
			}
			fmt.Println(_str)
		case _int, ok := <-ch2:
			if !ok {
				fmt.Println("CH2 Failed!")
			}
			fmt.Println(_int)
		default:
			// fmt.Println("随机")
		}
	}
}
