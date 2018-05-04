package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 8, 10}
	head := s1[0]
	tail := s1[len(s1)-1]
	out_head := s1[1:]
	out_tail := out_head[:len(out_head)-1]
	fmt.Println(head, tail)
	fmt.Println(out_head, out_tail)
}
