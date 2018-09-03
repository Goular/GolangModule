Go的运行时的当前设计，假定程序员自己负责检测何时终止一个goroutine以及何时终止该程序。 可以通过调用os.Exit或从main()函数的返回来以正常方式终止程序。而有时候我们需要的是使程序阻塞在这一行。

使用sync.WaitGroup
	一直等待直到WaitGroup等于0

		package main

		import "sync"

		func main() {
			var wg sync.WaitGroup
			wg.Add(1)
			wg.Wait()
		}

空select
	select{}是一个没有任何case的select，它会一直阻塞

		package main

		func main() {
			select{}
		}
		
死循环
	虽然能阻塞，但会100%占用一个cpu。不建议使用

		package main

		func main() {
			for {}
		}
		
用sync.Mutex
	一个已经锁了的锁，再锁一次会一直阻塞，这个不建议使用

		package main

		import "sync"

		func main() {
			var m sync.Mutex
			m.Lock()
			m.Lock()
		}
		
os.Signal
	系统信号量，在go里面也是个channel，在收到特定的消息之前一直阻塞

		package main

		import (
			"os"
			"syscall"
			"os/signal"
		)

		func main() {
			sig := make(chan os.Signal, 2)
			signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
			<-sig
		}
		
空channel或者nil channel
	channel会一直阻塞直到收到消息，nil channel永远阻塞。

		package main

		func main() {
			c := make(chan struct{})
			<-c
		}
		package main

		func main() {
			var c chan struct{} //nil channel
			<-c
		}
总结
	注意上面写的的代码大部分不能直接运行，都会panic，提示“all goroutines are asleep - deadlock!”，因为go的runtime会检查你所有的goroutine都卡住了， 没有一个要执行。你可以在阻塞代码前面加上一个或多个你自己业务逻辑的goroutine，这样就不会deadlock了。

参考
	http://pliutau.com/different-ways-to-block-go-runtime-forever/