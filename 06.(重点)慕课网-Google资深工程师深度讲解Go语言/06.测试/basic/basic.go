package basic

import (
	"fmt"
	"math"
)

func Triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 需要测试的方法
func CalcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 寻找不重复字符串 --  国际版
func LengthOfNonRepeatingSubStr(s string) int {
	// 创建map记录次数
	//lastOccurred := make(map[rune]int)
	// 使用更好的数据结构缓存数据
	lastOccurred := make([]int, 0xfffff)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		// fmt.Println("start::", start)
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
