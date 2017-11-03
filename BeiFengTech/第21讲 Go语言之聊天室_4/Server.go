package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)

var onlineConns = make(map[string]net.Conn)//存储客户端链接映射 key为链接ip:port value为连接对象conn
var messageQueue = make(chan string, 1000)//消息队列 带缓冲的buf

var quitChan = make(chan bool)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error:%s", err.Error())
		//系统退出，0代表正常退出，非零代表异常退出
		os.Exit(1)
	}
}

//添加连接的处理方法，一定不能放到主协程执行耗时内容，这样for的连续等待连接就没用了，而且会
//耗费系统资源,拖慢运行的时间
func ProcessInfo(conn net.Conn) { //这个方法属于生产者，因为将数据变为消息队列
	//由于是在协程工作，所以可以进行耗时操作
	buf := make([]byte, 1024)
	defer conn.Close();
	for {
		//第一个参数为缓冲的数量大小
		numOfBytes, err := conn.Read(buf);

		//由于一直轮询然,因为数据不是都是一直发的，可能就发一次，之后没数据的话就报错，后拿不到数据就报错，所以会一直出现err，这样就执行不了
		//CheckError(err)
		if err != nil {
			break
		}

		if numOfBytes != 0 {
			//同时返回对方的IP地址和端口
			//remoteAddr := conn.RemoteAddr()
			/*	结尾buf[0:numOfBytes]的原因是：numOfBytes是指接收的字节数 如果只用string(buf)
				可能会导致接收字符串中有其他之前接收的字符
			*/
			message := string(buf[:numOfBytes])
			messageQueue <- message
		}
	}
}

//消费者协程
func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			//对消息进行处理
			doProcessMessage(message)
		case <-quitChan:
			//直接退出循环，不再工作
			break
		}
	}
}

//对消息进行处理
func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		sendMessage := contents[1]

		addr = strings.Trim(addr, " ")

		//通过addr查看是否有链接对象
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure!")
			}
		}else{
			fmt.Println("找不到"+addr)
		}
	}
}

func main() {
	//设置监听端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	//defer即使是异常也会执行的
	defer listen_socket.Close()

	fmt.Println("Server is Waiting...")

	go ConsumeMessage()

	//由于访问人很多，所以我们进行死循环
	for {
		//如果有人连接，就进行接收与监听
		conn, err := listen_socket.Accept()
		CheckError(err)

		//判断访问者，同时将新的放着保存
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		//利用IP地址和端口的组合作为key值
		onlineConns[addr] = conn
		for key := range onlineConns {
			fmt.Println(key)
		}
		go ProcessInfo(conn)
	}
}
