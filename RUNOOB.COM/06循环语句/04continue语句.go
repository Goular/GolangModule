package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		if a == 15 {
			//跳过此次循环
			a += 1
			continue
		}
		fmt.Printf("a的值为:%d\n", a);
		a++
	}
}
