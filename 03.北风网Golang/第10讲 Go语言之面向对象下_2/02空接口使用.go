package main

import "fmt"

/**
	空接口的使用
 */
func main() {
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = 2.345
	var v4 interface{} = make(map[string]string)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}
