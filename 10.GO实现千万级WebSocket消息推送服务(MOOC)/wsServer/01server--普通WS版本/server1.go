package main

import (
	"net/http"
	"github.com/gorilla/websocket"
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
	 //w.Write([]byte("Hello World!")) // -- 这个是HTTP的应答
	// 更换成WebSocket的服务处理方法
	var (
		conn *websocket.Conn // 获得为长连接的conn实例
		err  error           // 异常处理
		data []byte
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	// websocket.Conn
	for {
		// 数据类型: Text，Binary -- 接收信息
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		// 返回信息
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:8888", nil)
}
