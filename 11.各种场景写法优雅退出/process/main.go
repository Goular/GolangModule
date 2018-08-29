package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

// 业务功能 ::: 优雅退出go守护进程

//Go中的Signal发送和处理
//
//golang中对信号的处理主要使用os/signal包中的两个方法：
//notify方法用来监听收到的信号
//stop方法用来取消监听

//SIGHUP	1	Term	终端控制进程结束(终端连接断开)
//SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
//SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
//SIGILL	4	Core	非法指令(程序错误、试图执行数据段、栈溢出等)
//SIGABRT	6	Core	调用abort函数触发
//SIGFPE	8	Core	算术运行错误(浮点运算错误、除数为零等)
//SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)
//SIGSEGV	11	Core	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
//SIGPIPE	13	Term	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
//SIGALRM	14	Term	时钟定时信号
//SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
//SIGUSR1	30,10,16	Term	用户保留
//SIGUSR2	31,12,17	Term	用户保留
//SIGCHLD	20,17,18	Ign	子进程结束(由父进程接收)
//SIGCONT	19,18,25	Cont	继续执行已经停止的进程(不能被阻塞)
//SIGSTOP	17,19,23	Stop	停止进程(不能被捕获、阻塞或忽略)
//SIGTSTP	18,20,24	Stop	停止进程(可以被捕获、阻塞或忽略)
//SIGTTIN	21,21,26	Stop	后台程序从终端中读取数据时触发
//SIGTTOU	22,22,27	Stop	后台程序向终端中写数据时触发

// 测试过程记录
//kill -USR1 pid 输出
//usr1 user defined signal 1
//
//kill -USR2 pid
//usr2 user defined signal 2
//
//kill pid
//退出 terminated
//开始退出...
//执行清理...
//结束退出...

// 优雅退出go守护进程
func main() {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	// todo: linux系统使用
	// signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
				// todo: SIGUSR1/SIGUSR2为linux系统使用
				//case syscall.SIGUSR1:
				//	fmt.Println("usr1", s)
				//case syscall.SIGUSR2:
				//	fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func ExitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
	// log.Fatal方法执行的是os.Exit(1)
	// log.Fatal("")
}
