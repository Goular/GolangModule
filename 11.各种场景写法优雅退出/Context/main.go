package main

import (
	"time"
	"fmt"
	"context"
)

// Context -- https://www.jianshu.com/p/dcbd87eb1a3f 这个教程讲的非常好

// chan+select的方式，是比较优雅的结束一个goroutine的方式，
// 不过这种方式也有局限性，如果有很多goroutine都需要控制结束怎么办呢？
// 如果这些goroutine又衍生了其他更多的goroutine怎么办呢？
// 如果一层层的无穷尽的goroutine呢？这就非常复杂了，
// 即使我们定义很多chan也很难解决这个问题，
// 因为goroutine的关系链就导致了这种场景非常复杂

// 上面说的这种场景是存在的，比如一个网络请求Request，
// 每个Request都需要开启一个goroutine做一些事情，
// 这些goroutine又可能会开启其他的goroutine。
// 所以我们需要一种可以跟踪goroutine的方案，才可以达到控制他们的目的，
// 这就是Go语言为我们提供的Context，称之为上下文非常贴切，它就是goroutine的上下文。

// context.Background() 返回一个空的Context，
// 这个空的Context一般用于整个Context树的根节点。
// 然后我们使用context.WithCancel(parent)函数，
// 创建一个可取消的子Context，然后当作参数传给goroutine使用，
// 这样就可以使用这个子Context跟踪这个goroutine。
//在goroutine中，使用select调用<-ctx.Done()判断是否要结束，
// 如果接受到值的话，就可以返回结束goroutine了；如果接收不到，就会继续进行监控。
//那么是如何发送结束指令的呢？这就是示例中的cancel函数啦，
// 它是我们调用context.WithCancel(parent)函数生成子Context的时候返回的，
// 第二个返回值就是这个取消函数，它是CancelFunc类型的。
// 我们调用它就可以发出取消指令，然后我们的监控goroutine就会收到信号，就会返回结束。
// 简单使用
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done():
//				fmt.Println("监控退出，停止了...")
//				return
//			default:
//				fmt.Println("goroutine监控中...")
//				time.Sleep(2 * time.Second)
//			}
//		}
//	}(ctx)
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("可以了，通知监控停止")
//	cancel()
//	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
//	time.Sleep(5 * time.Second)
//
//}

// Context控制多个goroutine
// 用于解决一个chan仅仅会被一个接收的尴尬问题
// 示例中启动了3个监控goroutine进行不断的监控，每一个都使用了Context进行跟踪，
// 当我们使用cancel函数通知取消时，这3个goroutine都会被结束。这就是Context的控制能力，
// 它就像一个控制器一样，按下开关后，所有基于这个Context或者衍生的子Context都会收到通知，
// 这时就可以进行清理操作了，
// 最终释放goroutine，这就优雅的解决了goroutine启动后不可控的问题。

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	go watch(ctx, "【监控1】")
//	go watch(ctx, "【监控2】")
//	go watch(ctx, "【监控3】")
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("可以了，通知监控停止")
//	cancel()
//	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
//	time.Sleep(5 * time.Second)
//}
//
//func watch(ctx context.Context, name string) {
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println(name, "监控退出，停止了...")
//			return
//		default:
//			fmt.Println(name, "goroutine监控中...")
//			time.Sleep(2 * time.Second)
//		}
//	}
//}

// Context的四个方法
//type Context interface {
//	Deadline() (deadline time.Time, ok bool)
//	Done() <-chan struct{}
//	Err() error
//	Value(key interface{}) interface{}
//}

// Deadline方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。
// Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。
// Err方法返回取消的错误原因，因为什么Context被取消。
// Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。
// 以上四个方法中常用的就是Done了，如果Context取消的时候，我们就可以得到一个关闭的chan，关闭的chan是可以读取的，所以只要可以读取的时候，就意味着收到Context取消的信号了，以下是这个方法的经典用法。

// Done的经典用法,同时可以返回Err异常
//func Stream(ctx context.Context, out chan<- Value) error {
//	for {
//		v, err := DoSomething(ctx)
//		if err != nil {
//			return err
//		}
//		select {
//		case <-ctx.Done():
//			return ctx.Err()
//		case out <- v:
//		}
//	}
//}

// todo: Context接口实现类
// 一个是Background，主要用于main函数、初始化以及测试代码中，
// 作为Context这个树结构的最顶层的Context，也就是根Context。

// todo :（大重点，一般什么都不用，或者要传递一个空context的时候使用这个，不要传递nil作为实参）一个是TODO,它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，
// 可以使用这个。

// 他们两个本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

// 由于零值的根Context基本没实际作用，所以为了达到各种各样的目的，我们需要使用Context的继承衍生
//func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
//func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
//func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
//func WithValue(parent Context, key, val interface{}) Context

//WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
//WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。
//WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到，后面我们会专门讲。

// WithValue 使用例子

// 在前面的例子，我们通过传递参数的方式，把name的值传递给监控函数。
//在这个例子里，我们实现一样的效果，但是通过的是Context的Value的方式。
//我们可以使用context.WithValue方法附加一对K-V的键值对，这里Key必须是等价性的，
//也就是具有可比性；Value值要是线程安全的。
//这样我们就生成了一个新的Context，这个新的Context带有这个键值对，
//在使用的时候，可以通过Value方法读取ctx.Value(key)。
//记住，使用WithValue传值，一般是必须的值，不要什么值都传递。

var key string="name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx:=context.WithValue(ctx,key,"【监控1】")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key),"监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key),"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

//Context 使用原则
//不要把Context放在结构体中，要以参数的方式传递
//以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
// todo: (大重点) 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
//Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
//Context是县城安全的，可以放心的在多个goroutine中传递
