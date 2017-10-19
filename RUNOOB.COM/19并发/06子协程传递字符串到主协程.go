package main

import (
	"fmt"
	"bytes"
	"strconv"
	"time"
)

func hello(xinDao chan string) {
	a := "I think I will use Golang In My Feture Project"
	//xinDao <- a

	for i := 0; i < 15; i++ {
		var buf bytes.Buffer
		buf.WriteString(a)
		buf.WriteString(strconv.Itoa(i))
		xinDao <- buf.String()
		time.Sleep(1 * time.Second)
	}
	xinDao <- "stop"
}

func main() {
	fmt.Println("创建子协程")
	xinDao := make(chan string)
	go hello(xinDao)
	fmt.Println("等待子协程信息的接收")
	for true {
		b := <-xinDao
		if b != "stop" {
			fmt.Println(b)
		} else {
			break;
		}
	}
	fmt.Println("主协程结束")
}
