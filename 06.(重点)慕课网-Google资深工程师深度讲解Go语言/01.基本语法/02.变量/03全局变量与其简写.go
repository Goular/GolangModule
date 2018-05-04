package main

import "fmt"

var (
	a int = 0
	b string
	c bool = true
)

func main() {
	fmt.Printf("%d %q %b", a, b, c)
}
