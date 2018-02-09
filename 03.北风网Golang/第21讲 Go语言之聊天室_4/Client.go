package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func MessageSend(conn net.Conn) {
	var input string
	for {
		//接收系统标准输入
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		//如果客户端输入exit 表示要结束连接
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			os.Exit(0)
			break
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client connect failure: " + err.Error())
			break
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer conn.Close()

	//开启消息发送协程
	go MessageSend(conn)

	//主协程负责接收消息
	buf := make([]byte, 1024)
	for {
		numOfBytes, err := conn.Read(buf)
		//CheckError(err)

		if err != nil{
			conn.Close()
			fmt.Println("Client connect failure :"+err.Error())
			break
		}

		/*结尾buf[0:numOfBytes]的原因是：numOfBytes是指接收的字节数 如果只用string(buf)
		可能会导致接收字符串中有其他之前接收的字符
		*/
		fmt.Println("receive server message content:" + string(buf[0:numOfBytes]))
	}

	fmt.Println("Client program end!")
}
