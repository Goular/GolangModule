package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return -1, nil
	}
}

func main() {
	_, err := Sqrt(-1)
	if (err != nil) {
		fmt.Println(err);
	}
}
