package main

import "fmt"

func main() {
	// 定义数组
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1::", s1)
	//fmt.Println(s2[4]) //越界
	fmt.Println("s2::", s2)
	fmt.Println("s1:len::", len(s1), "  cap::", cap(s1))
	fmt.Println("s2:len::", len(s2), "  cap::", cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	// 由于s4,s5的填充，所以我们的内容就会产生变化，arr会被新的内存分配赋值，
	fmt.Println(arr)
}
