package _2_RUNOOB例子代码

import (
	"fmt"
	"time"
)

func consumer(cname string, ch chan int) {
	for i := range ch {
		fmt.Println("consumer--", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("Producer--", pname, ":", i)
		ch <- i
		fmt.Println("阻塞",pname,i)
	}
}

func main() {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	data := make(chan int)
	go producer("生产者1", data)
	go producer("生产者2", data)
	go consumer("消费者1", data)
	go consumer("消费者2", data)

	time.Sleep(10 * time.Second)
	defer close(data);
	time.Sleep(10 * time.Second)
}
