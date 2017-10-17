package main

import (
	"fmt"
)

//第一种
var a1,a2 int
//外部只能进行声明，赋值的话会报错
//a1 = 11
//a2 = 12

//第二种
var b1,b2 = 21,22

//第四种
var(
	d1 int
	d2 bool = true
)

func main() {
	a1 = 11
	a2 = 12

	d1 = 41

	//第三种
	c1,c2 := 31,32

	fmt.Println(a1,a2)
	fmt.Println(b1,b2)
	fmt.Println(c1,c2)
	fmt.Println(d1,d2)
}

