package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	//linux crontab
	//秒粒度配置(2018-2099)

	//每隔5分钟执行1次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	//0,6,12,18,48

	//当前时间
	now = time.Now()

	//计算下次调度时间
	nextTime = expr.Next(now)
	fmt.Println(nextTime)

	//等待这个定时器超时
	//time.AfterFunc(nextTime.Sub(now), func() {
	//	fmt.Println("被调度了",nextTime)
	//})

	time.Sleep(5 * time.Second)
}
