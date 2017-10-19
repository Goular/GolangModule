package main

import "fmt"

//显式定义结构体
type Book struct {
	name string
	number int32
}


func main() {
	kitchen := Book{"诺曼底",1125}
	fmt.Println(kitchen.name)
}
