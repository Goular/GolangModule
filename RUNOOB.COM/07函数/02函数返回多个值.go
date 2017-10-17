package main

import (
	"fmt"
)

//参数返回并返回多参数
func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}
