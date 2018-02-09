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

	fmt.Println("Server is Waiting...")

	//由于访问人很多，所以我们进行死循环
	for {
		//如果有人连接，就进行接收与监听
		conn, err := listen_socket.Accept()
		CheckError(err)
		go ProcessInfo(conn)
	}
}

//添加连接的处理方法，一定不能放到主协程执行耗时内容，这样for的连续等待连接就没用了，而且会
//耗费系统资源,拖慢运行的时间
func ProcessInfo(conn net.Conn) {
	//由于是在协程工作，所以可以进行耗时操作
	buf := make([]byte, 1024)
	defer conn.Close();
	for {
		//第一个参数为缓冲的数量大小
		numOfBytes, err := conn.Read(buf);

		//由于一直轮询然,因为数据不是都是一直发的，可能就发一次，之后没数据的话就报错，后拿不到数据就报错，所以会一直出现err，这样就执行不了
		//CheckError(err)
		if err != nil{
			break
		}

		if numOfBytes != 0 {
			//同时返回对方的IP地址和端口
			remoteAddr := conn.RemoteAddr()
			fmt.Println(remoteAddr)
			fmt.Printf("Has Receive this Message: %s \n", string(buf[:numOfBytes]))
		}
	}
}
