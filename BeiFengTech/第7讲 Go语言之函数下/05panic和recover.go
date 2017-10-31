package main

import "fmt"

/**
	遇到错误，首先会执行defer栈的内容，然后才出现错误内容的错误树，如果使用recover来捕获，那么就不会产生错误树，而知根据捕获错误来处理
 */
func main() {
	defer func() {
		error :=recover()
		if (error!=nil){
			fmt.Println(error)
		}
	}()
	defer func() {
		fmt.Println("end 222")
	}()
	createPanic()
	fmt.Println("operate end")
}

func createPanic()  {
	panic("错误内容产生")
}
