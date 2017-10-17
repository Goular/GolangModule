package main

import "fmt"

//使用指针进行交换数据的的引用传递例子
func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

/* 定义相互交换值的函数 */
func swap(a, b *int) {
	var temp int
	temp = *a /* 保存 x 的值 */
	*a = *b    /* 将 y 值赋给 x */
	*b = temp /* 将 temp 值赋给 y*/
}
