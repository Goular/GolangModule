package main

import "fmt"

func main() {
	var a int = 20
	var ptr *int

	ptr = &a

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ptr )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ptr )
}
