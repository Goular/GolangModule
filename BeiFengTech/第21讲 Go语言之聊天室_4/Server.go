package main

import (
	"net"
	"fmt"
	"os"
	"strings"
	"log"
)

//定义日志的位置
const (
	//记录错误日志的路径
	LOG_DIRECTORY = "./test.log"
)

var onlineConns = make(map[string]net.Conn) //存储客户端链接映射 key为链接ip:port value为连接对象conn
var messageQueue = make(chan string, 1000)  //消息队列 带缓冲的buf
var logFile *os.File
var logger *log.Logger

var quitChan = make(chan bool)

func CheckError(err error) {
	if err != nil {
		panic(err)
		//fmt.Println("Error:%s", err.Error())
		////系统退出，0代表正常退出，非零代表异常退出
		//os.Exit(1)
	}
}

//添加连接的处理方法，一定不能放到主协程执行耗时内容，这样for的连续等待连接就没用了，而且会
//耗费系统资源,拖慢运行的时间
func ProcessInfo(conn net.Conn) { //这个方法属于生产者，因为将数据变为消息队列
	//由于是在协程工作，所以可以进行耗时操作
	buf := make([]byte, 1024)
	defer func(conn net.Conn) {
		//查找连接的IP
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		//删除Map中关于相关IP的连接
		delete(onlineConns, addr)
		conn.Close();
		for key := range onlineConns {
			fmt.Println("now online conns:" + key)
		}
	}(conn) //采用匿名函数的方式 调用defer
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
		sendMessage := strings.Join(contents[1:], "#")//这么做是为了防止 消息体也含有"#"

		addr = strings.Trim(addr, " ")

		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure!")
			}
		}
	} else {
		//走到这里 说明客户端调用list命令 查看系统当前链接ip
		contents := strings.Split(message, "*")
		if strings.ToUpper(contents[1]) == "LIST" {
			var ips string = ""
			for i := range onlineConns {
				ips = ips + "|" + i
			}
			if conn, ok := onlineConns[contents[0]]; ok {
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("online conns send failure!")
				}
			}
		}
	}
}

func main() {
	//设置日志打开
	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("log file create failure!")
		os.Exit(-1)
	}
	defer logFile.Close()
	//利用go自带的log 将打开文件对象生成日志文件对象
	logger = log.New(logFile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)

	//设置监听端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	//defer即使是异常也会执行的
	defer listen_socket.Close()

	fmt.Println("Server is Waiting...")

	logger.Println("I am writing the logs...")

	//开启消费者协程
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
