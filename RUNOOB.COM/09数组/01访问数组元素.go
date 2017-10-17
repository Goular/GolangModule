package main

import "fmt"

func main() {
	//定义一个长度为10的整型数组
	var n [10]int
	var i, j int

	//为数组n初始化元素
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}
	//输出每个数组元素的值
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
