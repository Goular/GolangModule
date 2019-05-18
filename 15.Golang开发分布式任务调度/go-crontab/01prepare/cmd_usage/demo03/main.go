package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

// 结果结构体
type result struct {
	err    error
	output []byte
}

//  执行1个cmd, 让它在一个协程里去执行, 让它执行2秒: sleep 2; echo hello;
// 1秒的时候, 我们杀死cmd
func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)
	// 创建一个channel队列
	resultChan = make(chan *result, 1000)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "E:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2;echo hello;")
		// 执行任务，捕获任务
		output, err = cmd.CombinedOutput()
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	// 继续往下面走
	time.Sleep(1 * time.Second)
	// 通知上下文关闭channel
	cancelFunc()
	// 等待协程，打印内容
	res = <-resultChan
	// 打印任务结果
	fmt.Println(res.err, string(res.output))
}
