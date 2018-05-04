package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func main() {
	euler()
}
