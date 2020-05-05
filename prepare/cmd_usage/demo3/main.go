package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	//执行1个cmd，让它在一个携程李执行，让他执行2秒 sleep 2 echo hello

	//一秒的时候 我们杀死他
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	resultChan = make(chan *result, 1000)
	//context chan byte
	//cancelFunc close（chan byte）

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)

		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2,echo hello")
		//select {case <- ctx.Done():}
		//kill pif 进程ID 杀死]

		//执行任务，捕获输出
		output, err = cmd.CombinedOutput()

		resultChan <- &result{
			err:    err,
			output: output,
		}

	}()

	//继续往下走
	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	//在main协程，等待
	res = <-resultChan

	fmt.Println(res.err, string(res.output))
}
