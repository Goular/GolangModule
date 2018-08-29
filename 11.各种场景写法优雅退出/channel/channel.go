package main

import (
	"math/rand"
	"time"
	"log"
	"sync"
	"strconv"
)

// 优雅的关闭Golang Channel

// 这个例子中生产端和接受端都没有关闭消息数据的channel，channel在没有任何goroutine引用的时候会自行关闭，而不需要显示进行关闭。

// todo: Channel关闭原则
// 不要在消费端关闭channel，不要在有多个并行的生产者时对channel执行关闭操作。

// 也就是说应该只在[唯一的或者最后唯一剩下]的生产者协程中关闭channel，来通知消费者已经没有值可以继续读了。只要坚持这个原则，就可以确保向一个已经关闭的channel发送数据的情况不可能发生。

// once 关闭channel方法
//type MyChannel struct {
//	C    chan T
//	once sync.Once
//}
//
//func NewMyChannel() *MyChannel {
//	return &MyChannel{C: make(chan T)}
//}
//
//func (mc *MyChannel) SafeClose() {
//	mc.once.Do(func() {
//		close(mc.C)
//	})
//}


// 多个消费者，单个生产者的channel关闭
//func main() {
//	// 生成随机数
//	rand.Seed(time.Now().UnixNano())
//	log.SetFlags(0)
//	// 定义访问的参数
//	const MaxRandomNumber = 100000
//	const NumReceivers = 100
//
//	wgReceiver := sync.WaitGroup{}
//	wgReceiver.Add(NumReceivers)
//
//	dataCh := make(chan int, 100)
//
//	// Send
//	go func() {
//		for {
//			if value := rand.Intn(MaxRandomNumber); value == 0 {
//				close(dataCh)
//				return
//			} else {
//				dataCh <- value
//			}
//		}
//	}()
//
//	// Receive
//	for i := 0; i < NumReceivers; i++ {
//		go func() {
//			defer wgReceiver.Done()
//			for value := range dataCh {
//				log.Println(value)
//			}
//		}()
//	}
//
//	wgReceiver.Wait()
//}

//多个生产者，单个消费者
//生产者同时也是退出信号channel的接受者，退出信号channel仍然是由它的生产端关闭的，
// 所以这仍然没有违背channel关闭原则。值得注意的是，这个例子中生产端和接受端都没有关闭消息数据的channel，
// channel在没有任何goroutine引用的时候会自行关闭，而不需要显示进行关闭。
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	log.SetFlags(0)
//
//	// ...
//	const MaxRandomNumber = 100000
//	const NumSenders = 1000
//
//	wgReceivers := sync.WaitGroup{}
//	wgReceivers.Add(1)
//
//	// ...
//	dataCh := make(chan int, 100)
//	stopCh := make(chan struct{})
//	// stopCh is an additional signal channel.
//	// Its sender is the receiver of channel dataCh.
//	// Its reveivers are the senders of channel dataCh.
//
//	// senders
//	for i := 0; i < NumSenders; i++ {
//		go func() {
//			for {
//				// The first select here is to try to exit the goroutine
//				// as early as possible. In fact, it is not essential
//				// for this example, so it can be omitted.
//				//select {
//				//case <- stopCh:
//				//	return
//				//default:
//				//}
//
//				// Even if stopCh is closed, the first branch in the
//				// second select may be still not selected for some
//				// loops if the send to dataCh is also unblocked.
//				// But this is acceptable, so the first select
//				// can be omitted.
//				select {
//				case <- stopCh:
//					return
//				case dataCh <- rand.Intn(MaxRandomNumber):
//				}
//			}
//		}()
//	}
//
//	// the receiver
//	go func() {
//		defer wgReceivers.Done()
//
//		for value := range dataCh {
//			if value == MaxRandomNumber-1 {
//				// The receiver of the dataCh channel is
//				// also the sender of the stopCh cahnnel.
//				// It is safe to close the stop channel here.
//				close(stopCh)
//				return
//			}
//
//			log.Println(value)
//		}
//	}()
//
//	// ...
//	wgReceivers.Wait()
//}

// 多个生产者，多个消费者
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// The channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// Here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The first select here is to try to exit the goroutine
				// as early as possible. This select blocks with one
				// receive operation case and one default branches will
				// be optimized as a try-receive operation by the
				// official Go compiler.
				//select {
				//case <- stopCh:
				//	return
				//default:
				//}

				// Even if stopCh is closed, the first branch in the
				// second select may be still not selected for some
				// loops (and for ever in theory) if the send to
				// dataCh is also unblocked.
				// This is why the first select block is needed.
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// Same as the sender goroutine, the first select here
				// is to try to exit the goroutine as early as possible.
				//select {
				//case <- stopCh:
				//	return
				//default:
				//}

				// Even if stopCh is closed, the first branch in the
				// second select may be still not selected for some
				// loops (and for ever in theory) if the receive from
				// dataCh is also unblocked.
				// This is why the first select block is needed.
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						// The same trick is used to notify
						// the moderator to close the
						// additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
