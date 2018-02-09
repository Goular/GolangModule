package main

import "fmt"

var identifier []int

func main() {
	var slice1 []int = make([]int,5)
	slice2 := make([]string,6)
	fmt.Println(slice1)
	fmt.Println(slice2)

	s2 := []int{2,5,8}
	//s := []int{1,2,3}
	s1 := s2
	s3 := s2[0:2]
	s4 := s2[2:]
	s5 := s2[:5]
	s6 := s2[1:2]
	s7 := make([]int,5,9)
}
