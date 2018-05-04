package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//func ServeRpc(host string, service interface{}) error {
//	rpc.Register(service)
//	// 创建TCP连接
//	listener, err := net.Listen("tcp", host)
//	if err != nil {
//		return err
//	}
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Printf("accept error :%v", err)
//			continue
//		}
//		// 为了不阻塞tcp所以我们开一个协程进行处理
//		// 开了协程一般不会出错
//		go jsonrpc.ServeConn(conn)
//	}
//	return nil
//}
//
//func NewClient(host string) (*rpc.Client, error) {
//	conn, err := net.Dial("tcp", ":1234")
//	if err != nil {
//		return nil, err
//	}
//	return jsonrpc.NewClient(conn), nil
//}

func ServeRpc(
	host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	log.Printf("Listening on %s", host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
