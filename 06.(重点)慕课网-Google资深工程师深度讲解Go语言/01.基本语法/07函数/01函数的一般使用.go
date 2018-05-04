package main

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("不支持相关的操作符")
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

// 这种写法一般不用，因为函数的长度太多，起名字会让代码不好读
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func main() {
	fmt.Println(eval(3, 4, "+"))
	fmt.Println(div(13, 3))
}
