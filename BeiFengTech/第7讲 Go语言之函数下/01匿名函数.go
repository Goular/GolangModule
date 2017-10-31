package main

import "fmt"

func main() {
	type sum func()
	var y sum
	y = func() {
		fmt.Println("Hello World！12333")
	}
	y()

	x := func(a, b int) int {
		return a + b
	}
	fmt.Println(x(2,8))
}
