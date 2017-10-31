package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 2 + 4i
	fmt.Print(x + x)
	fmt.Println()
	fmt.Print(reflect.TypeOf(x)) //获取变量X的类型
}
