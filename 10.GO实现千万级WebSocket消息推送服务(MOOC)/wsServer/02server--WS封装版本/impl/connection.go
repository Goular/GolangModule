package impl

// select 特性
// 如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
// 如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
// 如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

import (
	"github.com/gorilla/websocket"
	"sync"
	"errors"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	isClosed bool       // 判断closeChan是否关闭了,但是由于多线程访问的问题，所以我们需要添加一把锁
	mutex    sync.Mutex // 用于解决同步数据冲突问题
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000), // 容量为1000，超过的内容就不处理了
		outChan:   make(chan []byte, 1000), // 容量为1000，超过的内容就不处理了
		closeChan: make(chan byte, 1),
	}
	// 启动读协程
	go conn.readLoop()

	// 启动写协程
	go conn.writeLoop()

	return
}

// API封装
func (conn *Connection) ReadMessage() (data []byte, err error) {
	// 这个位置也会出现底层网络出错，导致chan阻塞的问题
	// data = <-conn.inChan

	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("Connection is Closed.")
	}
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	// 这个位置也会出现阻塞，导致chan阻塞的问题
	// conn.outChan <- data

	select {
	case <-conn.closeChan:
		err = errors.New("Connection is Closed.")
	case conn.outChan<-data:
	}
	return
}

func (conn *Connection) Close() {
	// 该方法线程安全，可重入的Close，任意线程调用都不会有问题
	conn.wsConn.Close()

	// todo:多次执行改行代码会报错panic，所以我们需要保证这一行代码只执行一次
	// close(conn.closeChan)

	// 下面这段代码能够保证线程安全且可以重复执行
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		// 一直读取长连接的信息，读到就投递到channel
		// 存在阻塞的问题，阻塞时会等待inChan有空闲的时候再进行写入channel，
		// 需要注意问题，就是如果我们的inChan阻塞了，不继续走for，此时，如果写chan出现问题，这样连接conn.close执行了，这样写协程关闭了连接,这个阻塞永远没有办法解开，此时会出现泄漏协程的问题
		// todo:为了能够解决阻塞带来的协程泄漏，out of control的问题，我们采用多路复用select
		// conn.inChan <- data

		select {
		case <-conn.closeChan: // closeChan关闭的时候会触发这个内容
			goto ERR
		case conn.inChan <- data:
		}
	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		// <-conn.outChan如果一直没有触发也会被阻塞，所以也是用多路复用select来解决
		// data = <-conn.outChan
		select {
		case <-conn.closeChan: // 关闭当前的写协程
			goto ERR
		case data = <-conn.outChan: // 使用select防止这个语句阻塞
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
