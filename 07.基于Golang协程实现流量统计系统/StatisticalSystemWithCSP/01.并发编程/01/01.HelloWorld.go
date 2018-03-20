package main

import "fmt"

var message = make(chan string)

func sample() {
	message <- "Hello Imooc!"
}

func sample2() {
	message <- "I'm Sample2"
}

func main() {
	go sample()
	go sample2()
	fmt.Println("Hello World！")
	fmt.Println(<-message)
	fmt.Println(<-message)
}
