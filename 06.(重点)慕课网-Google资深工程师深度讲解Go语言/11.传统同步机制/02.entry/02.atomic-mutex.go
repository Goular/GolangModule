package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("啦啦啦啦，逻辑代码")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
	fmt.Println("啦啦啦啦，逻辑代码")
}

func (a *atomicInt) get() int {
	fmt.Println("啦啦啦啦，逻辑代码")
	a.lock.Lock()
	// 这个锁在函数执行完的最后才执行，可能会影响后面的代码，因为阻塞了，所以我们采用匿名函数的方式并执行的方式解决问题
	defer a.lock.Unlock()
	fmt.Println("啦啦啦啦，逻辑代码")
	return a.value
}

// 运行后存在冲突
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
