package main

import "fmt"


func main() {
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	//会自动添加并赋值给下面的b2,c2
	const (
		a2 = iota
		b2
		c2
	)

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}
