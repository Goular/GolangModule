package main

import "fmt"

func main() {
	// 方法1
	s1 := []int{2, 4, 6, 8, 10}
	printSlice(s1)
	// 方法2
	s2 := make([]int, 16)
	printSlice(s2)
	// 方法3
	s3 := make([]int, 10, 32)
	printSlice(s3)
}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}
