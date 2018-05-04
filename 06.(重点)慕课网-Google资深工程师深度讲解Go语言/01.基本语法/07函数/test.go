package main

import "fmt"

const (
	// 字节顺序
	BYTE_ORDER_ABCD = "1"
	BYTE_ORDER_CDAB = "2"
	BYTE_ORDER_BADC = "3"
	BYTE_ORDER_DCBA = "4"
)

func main() {
	// 原始数据
	bytes := []byte("abcd")
	fmt.Println(exchangeBytesOrder(bytes, "4"))
}
func exchangeBytesOrder(datas []byte, bit string) []byte {
	switch bit {
	case BYTE_ORDER_ABCD:
		return []byte{datas[0], datas[1], datas[2], datas[3]}
	case BYTE_ORDER_CDAB:
		return []byte{datas[2], datas[3], datas[0], datas[1]}
	case BYTE_ORDER_BADC:
		return []byte{datas[1], datas[0], datas[3], datas[2]}
	case BYTE_ORDER_DCBA:
		return []byte{datas[3], datas[2], datas[1], datas[0]}
	default:
		return []byte{}
	}
}
