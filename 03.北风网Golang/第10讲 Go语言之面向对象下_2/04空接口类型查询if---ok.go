package main

import (
	"fmt"
)

func main() {
	var v1 interface{}
	v1 = "123"
	//v1 = 11223
	if value, ok := v1.(int); ok {
		fmt.Println(value, ok)
	} else {
		fmt.Println("暂时不属于此类型!")
	}
}
