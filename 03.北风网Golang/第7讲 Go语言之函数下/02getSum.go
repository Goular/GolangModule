package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6}
	sum := getSum(num)
	fmt.Println(sum)
}

func getSum(num []int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}
