package main

import "fmt"

//当创建函数时，你定义了函数需要做什么，通过调用改函数来执行指定任务。
func main() {
	//定义局部数据
	a, b := 100, 200
	var ret int
	ret = max(a, b);
	fmt.Printf("最大值是 : %d\n", ret)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}
