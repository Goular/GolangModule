package main

import "fmt"

func main() {
	/* main 函数中声明局部变量 */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中 a = %d\n",  a);
	c = sum( a, b);
	fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n",  a);
	fmt.Printf("sum() 函数中 b = %d\n",  b);

	return a + b;
}
