package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 interface{}
	//v1 = "123"
	v1 = 11223
	switch v1.(type) {
	case int:
		fmt.Println("Int")
	case int8:
		fmt.Println("Int8")
	case int32:
		fmt.Println("Int32")
	case float32:
		fmt.Println("Float32")
	case string:
		fmt.Println("字符串")
	default:
		fmt.Println("暂无相关类型")
	}
	fmt.Println(reflect.TypeOf(v1))
}
