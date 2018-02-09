package main

import (
	"fmt"
)

func main() {
	for i:=1;i<=5;i++{
		defer fmt.Println(i)
	}
	fmt.Println("Behind To 223344")
}
