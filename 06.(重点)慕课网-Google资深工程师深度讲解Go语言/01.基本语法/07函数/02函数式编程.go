package main

import (
	"reflect"
	"runtime"
	"fmt"
	"math"
)

func apply(op func(int, int) int, a, b int) int {
	// 获取指针，用于获取函数名称
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 累加算法
func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

var p1 func(int, int) int

func main() {
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 4, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
