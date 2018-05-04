package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

// 切片其实就是数组的视图
func main() {
	// 数组定义
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 数组内容转切片
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	s1 := arr[2:6]
	s2 := arr[:]

	// 切片的变量是一个指针
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
}
