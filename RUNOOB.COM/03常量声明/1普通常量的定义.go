package main

import "fmt"

const (
	a1 = true
	b1 = "天鹅湖"
	c1 = 8.28
)

func main() {
	const LENGTH int = 100
	const WIDTH int = 200
	var area int
	const a, b, c = 1, false, "加油"

	area = LENGTH * WIDTH;
	fmt.Printf("面积为:%d\n", area)

	println()
	println(a, b, c)
	println(a1, b1, c1)
}
