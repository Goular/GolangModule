package main

import (
	"fmt"
	"time"
)

func test_Routine()  {
	fmt.Println("This is one routine!!!")
}

func Add(x,y int)  {
	z:=x+y
	fmt.Println(z)
}

func main() {
	//简单使用协程
	//go test_Routine()


	for i:=1;i<10;i++{
		go Add(i,i)
	}

	time.Sleep(2)
}
