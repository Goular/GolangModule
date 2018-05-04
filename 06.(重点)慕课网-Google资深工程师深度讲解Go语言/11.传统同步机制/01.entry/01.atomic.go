package main

import (
	"fmt"
	"time"
)

// 读写容易产生冲突
type atomicInt int

func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

// 运行后存在冲突
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
