package main

import (
	"net"
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error:%s", err.Error())
		//系统退出，0代表正常退出，非零代表异常退出
		os.Exit(1)
	}
}

func main() {
	//设置监听端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	//defer即使是异常也会执行的
	defer listen_socket.Close()

}
