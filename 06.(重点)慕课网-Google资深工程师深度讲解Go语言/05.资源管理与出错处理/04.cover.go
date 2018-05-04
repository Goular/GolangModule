package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err)
		} else {
			// 如果不是error，直接执行panic
			panic(fmt.Sprintf("I don't know what to do ：%v", r))
		}
	}()
	// panic(errors.New("This is An Error"))
	panic(123)
	b := 0
	fmt.Println(5 / b)
}

func main() {
	tryRecover()
}
