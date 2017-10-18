package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("欢迎来到Golang协程的世界")
}

//写法一
//由于主协程直接执行下一条打印后没有代码执行，直接退出，所以子协程的返回输出因主协程的结束而结束，没有输出
//func main() {
//	go hello()
//	fmt.Println("Main Function")
//}

//写法二
//创建执行hello子协程后，休眠1s，然后继续执行主协程，这样子协程执行完返回的内容能够成功输出
//注意:在主协程中使用 Sleep 函数等待其他协程结束的方法是不正规的
func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Main Function")
}
