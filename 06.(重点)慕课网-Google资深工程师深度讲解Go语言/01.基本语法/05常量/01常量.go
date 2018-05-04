package main

import (
	"fmt"
	"math"
)

const (
	ab     = "123"
	ac     = 123
	ae     = true
	ar, ap = true, 123
)

func main() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(fileName, c)
}
