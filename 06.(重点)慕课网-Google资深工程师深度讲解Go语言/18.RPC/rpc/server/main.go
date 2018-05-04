package main

import (
	"net/rpc"

	"log"
	"net"

	"net/rpc/jsonrpc"

	"baseLearn/18.RPC/rpc"
)

func main() {
	rpc.Register(rpcDemo.DemoService{})
	// 创建TCP连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error :%v", err)
			continue
		}
		// 为了不阻塞tcp所以我们开一个协程进行处理
		go jsonrpc.ServeConn(conn)
	}
}
