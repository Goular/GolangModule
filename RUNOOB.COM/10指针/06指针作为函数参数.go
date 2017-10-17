package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )
}

func swap(a, b *int) {
	var temp int
	temp = *a /* 保存 x 地址的值 */
	*a = *b   /* 将 y 赋值给 x */
	*b = temp /* 将 temp 赋值给 y */
}
