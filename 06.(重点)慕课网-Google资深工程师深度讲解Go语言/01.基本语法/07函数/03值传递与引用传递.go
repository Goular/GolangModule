package main

import "fmt"

// 使用指针值传递做交换
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 使用返回值做参数交换（推荐使用）
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swap2(5, 6))
}
