package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	// 创建map记录次数
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		fmt.Println("start::", start)
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
}
