package demo1

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

var (
	expr     *cronexpr.Expression // cron的任务的对象
	err      error
	now      time.Time // 当前时间
	nextTime time.Time // 下一次运行的时间
	flag     = make(chan struct{})
)

func main() {
	// linux crontab
	// 秒粒度, 年配置(2018-2099)
	// 哪一分钟（0-59），哪小时（0-23），哪天（1-31），哪月（1-12），星期几（0-6）

	// 定义5分钟运行一次 每隔5分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	// 获取当前时间
	now = time.Now()
	//下一次调度的时间
	nextTime = expr.Next(now)

	fmt.Println(now)
	fmt.Println(nextTime)

	// 计算两者的时间差，等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("任务被调度了:", nextTime)
	})

	select {}
}
