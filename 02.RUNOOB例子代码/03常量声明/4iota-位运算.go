package main

import "fmt"

const (
	i = 1<<iota
	j
	k
	l
)

const (
	i2 = 1>>iota
	j2
	k2
	l2
)

func main() {
	//每次定义的往左移动一位,iota在每次定义的时候都增加1
	fmt.Println("i=",i) //1 -- 1>>1
	fmt.Println("j=",j) //2 -- 1>>2
	fmt.Println("k=",k) //4 -- 1>>3
	fmt.Println("l=",l) //8 -- 1>>4

	fmt.Println("i2=",i2) //1
	fmt.Println("j2=",j2) //2
	fmt.Println("k2=",k2) //4
	fmt.Println("l2=",l2) //8
}
