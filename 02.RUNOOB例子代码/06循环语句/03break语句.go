package main

import (
	"fmt"
)

func main() {
	//Switch不使用break，现在使用break只有for了
	var a int = 10
	//for循环
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			//使用break语句跳出循环
			break
		}
	}
}
