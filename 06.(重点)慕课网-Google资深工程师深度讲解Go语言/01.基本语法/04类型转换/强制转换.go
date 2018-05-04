package main

import (
	"math"
	"fmt"
)

func main() {
	a, b := 3, 4
	result := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(result)

}
