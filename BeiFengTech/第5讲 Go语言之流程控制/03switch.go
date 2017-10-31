package main

import "fmt"

func main() {
	i:=3
	switch i {
	case 1,2:
		fmt.Printf("111")
	case 3:fallthrough //case穿透的使用
	case 4:
		fmt.Printf("222")
	default:
		fmt.Printf("000")

	}
}
