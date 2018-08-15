package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"wsServer/02server--WS封装版本/impl"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		// 由于浏览器访问地址一般会存在跨域，所以在检查是否允许跨域时，我们一般都会直接允许
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}
	// TODO: 下面代码用于测试线程安全的问题,这样也不会出现心跳包的问题
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
// TODO: 关闭连接操作
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:8888", nil)
}
