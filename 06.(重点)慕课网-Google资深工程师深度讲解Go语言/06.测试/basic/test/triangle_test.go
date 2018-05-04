package test

import (
	"testing"
	//. "baseLearn/06.测试/basic"
)

//func TestTriangle(t *testing.T) {
//	tests := []struct{ a, b, c int }{
//		{3, 4, 5},
//		{5, 12, 13},
//		{12, 35, 0},
//		{30000, 40000, 50000},
//	}
//	for _, tt := range tests {
//		if actual := CalcTriangle(tt.a, tt.b); actual != tt.c {
//			// 如果与C不相等的话
//			t.Errorf("CalcTriangle(%d,%d);Got %d;Expected %d", tt.a, tt.b, actual, tt.c)
//		}
//	}
//}

//fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
//fmt.Println(lengthOfNonRepeatingSubStr("bbbbbb"))
//fmt.Println(lengthOfNonRepeatingSubStr("pwwke2w"))
//fmt.Println(lengthOfNonRepeatingSubStr("统一奶茶一"))
//fmt.Println(lengthOfNonRepeatingSubStr("天南第一峰9958"))

// 寻找不重复字符串 --  国际版
func LengthOfNonRepeatingSubStr(s string) int {
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

// 最长不重复字符串数量
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// 普通的字符串
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwke2w", 4},
		// 中文支持的字符串
		{"统一奶茶一", 4},
		{"天南第一峰9958", 6},
		{"黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞", 8},
	}
	for _, tt := range tests {
		if actual := LengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("Got %d for input %s;"+"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞"
	answer := 8
	for i := 0; i < b.N; i++ {
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != answer {
			b.Errorf("Got %d for input %s expected %d", actual, s, answer)
		}
	}
}
