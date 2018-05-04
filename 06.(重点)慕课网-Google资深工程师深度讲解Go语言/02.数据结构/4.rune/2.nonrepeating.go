package main

import "fmt"

// 寻找不重复字符串 --  国际版
func lengthOfNonRepeatingSubStr(s string) int {
	// 创建map记录次数
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		// fmt.Println("start::", start)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwke2w"))
	fmt.Println(lengthOfNonRepeatingSubStr("统一奶茶一"))
	fmt.Println(lengthOfNonRepeatingSubStr("天南第一峰9958"))
}
