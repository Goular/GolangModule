package main

import (
	"fmt"
	"os/exec"
)

// 该main方法没有进行捕获输出
func main() {
	var (
		cmd *exec.Cmd
		err error
	)
	// todo：下面的指令针对的linux系统
	//cmd = exec.Command("/bin/bash", "-c", "echo 1;echo 2;")
	// todo:下面的指令针对的window系统
	cmd = exec.Command("E:\\cygwin64\\bin\\bash.exe", "-c", "echo 1;echo 3;")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
