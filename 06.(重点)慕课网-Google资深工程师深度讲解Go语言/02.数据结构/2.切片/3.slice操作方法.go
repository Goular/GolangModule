package main

import "fmt"

func main() {
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	fmt.Println("演示复制Slice")
	s1 := make([]int, 100)
	copy(s1, s)
	fmt.Println(s1)

	fmt.Println("删除Slice")
	s2 := append(s1[:3], s1[4:]...) // 省略符号用于将slice拆分为一个不定长的整型队列
	fmt.Println(s2)
}

func printSlice(s []int) {
	fmt.Println("len=%d,cap=%d\n", len(s), cap(s))
}
