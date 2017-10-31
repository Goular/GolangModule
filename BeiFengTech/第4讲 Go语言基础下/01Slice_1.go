package main

import "fmt"

func main() {
	x:= [...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(len(x))

	var y1 []int
	y1 = x[3:]
	fmt.Printf("%v\n", y1)

	y2 := x[:3]
	fmt.Printf("%v\n", y2)
}
