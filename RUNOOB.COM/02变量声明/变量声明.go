package main

import "fmt"

func main() {
	//第一种
	var a int
	a = 11
	fmt.Println("a:", a);
	//第二种
	var b  = 22
	fmt.Println("b:", b);
	//第三种(不能用在全局变量定义中)
	c := 33;
	fmt.Println("c:", c);
}
