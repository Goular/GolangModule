package main

import "fmt"

func main() {
	//a := 10
	//if a<20 {
	//	fmt.Printf("a 小于 20 \n")
	//}else{
	//	fmt.Printf("a 大于 20 \n")
	//}
	//fmt.Printf("a 的值是 %d \n",a)

	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
		}
	}
	fmt.Printf("a 值为 : %d\n", a );
	fmt.Printf("b 值为 : %d\n", b );
}
