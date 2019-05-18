package main

import (
	"fmt"
	"os/exec"
	"time"
)

// 该main方法没有进行捕获输出
func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)
	// 生成CMD
	cmd = exec.Command("E:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2;ls -l")
	// 执行命令，不获取子进程的输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1 * time.Second)
	// 捕获到内容
	fmt.Println(string(output))
}
