package main

import "fmt"

func swap(a,b int)(int,int)  {
	return b,a
}

func add(a *int)*int  {
	*a = *a+1
	return a
}

func main()  {
	a:=5
	add(&a)
	fmt.Println(a)
}
