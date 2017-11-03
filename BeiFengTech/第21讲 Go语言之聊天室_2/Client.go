package main

import (
	"fmt"
	"os"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error:%s", err.Error())
		//系统退出，0代表正常退出，非零代表异常退出
		os.Exit(1)
	}
}

func main() {
	//客户端要进行拨号
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer conn.Close()

	//进行写 操作
	conn.Write([]byte("Hello,I come from GuangZhou"))
	fmt.Println("Client sent the Message to Server")
}
