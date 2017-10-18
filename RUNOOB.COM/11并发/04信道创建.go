package main

import "fmt"

//创建信道
func main() {
	var a chan int
	if a == nil {
		fmt.Println("Channel a is nil,going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}
