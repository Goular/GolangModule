package main

import "fmt"

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 4, 5}

	//for 循环
	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为:%d\n", a)
	}

	fmt.Println("********")

	for a < b {
		a++
		fmt.Printf("a的值为:%d\n", a)
	}

	fmt.Println("********")

	for key, value := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", key, value)
	}
}
