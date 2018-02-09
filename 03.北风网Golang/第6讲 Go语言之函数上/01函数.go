package main

import "fmt"

func main() {
	x:=2
	x2 := &x
	println(x2)
	println(*x2)
	a:=4
	add(&a)
	fmt.Println(a)
}

func swap(a,b int)(int,int)  {
	return b,a
}

func add(a *int) *int {
	*a = *a +1
	return a
}
