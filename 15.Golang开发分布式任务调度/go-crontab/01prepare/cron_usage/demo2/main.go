package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

// 定义任务类
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob // key为任务的名称
	)
	scheduleTable = make(map[string]*CronJob)
	// 获取当前时间
	now = time.Now()
	// 1.定义两个crontab
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 将任务注册到调度表
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 将任务注册到调度表
	scheduleTable["job2"] = cronJob

	// 开启一个调度协程
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		for {
			now = time.Now()
			//fmt.Println("aaaaaa")
			for jobName, cronJob = range scheduleTable {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					go func(jobName string) {
						fmt.Println("执行:", jobName)
					}(jobName)
					// 计算下一次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间:", cronJob.nextTime)
				}
			}
			// 睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:

			}

		}
	}()
	select {

	}
}
