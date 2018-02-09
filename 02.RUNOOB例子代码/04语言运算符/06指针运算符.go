package main

import "fmt"

func main() {
	var a int = 11
	var b int32 = 12333
	var c float32 =8.5624
	var ptr *int
	var ptr2 *int32
	var ptr3 *float32


	/**
	 *	运算符实例
	 */
	fmt.Println("第1行 - a 变量类型为 = %T\n",a)
	fmt.Println("第2行 - a 变量类型为 = %T\n",b)
	fmt.Println("第3行 - a 变量类型为 = %T\n",c)

	/*
	 *	& 和 * 运算符实例
	 */
	ptr = &a
	ptr2 = &b
	ptr3 = &c
	fmt.Println("a的值为 %d\n",a)

	fmt.Println("*ptr的内存值为 %d\n",ptr)
	fmt.Println("ptr的值为 %d\n",*ptr)

	fmt.Println("*ptr2的内存值为 %d\n",ptr2)
	fmt.Println("ptr2的值为 %d\n",*ptr2)

	fmt.Println("*ptr3的内存值为 %d\n",ptr3)
	fmt.Println("ptr3的值为 %d\n",*ptr3)
}