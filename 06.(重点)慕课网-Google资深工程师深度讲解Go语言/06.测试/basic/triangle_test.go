package basic

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
	}
	for _, tt := range tests {
		if actual := LengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("Got %d for input %s;"+"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	answer := 8
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != answer {
			b.Errorf("Got %d for input %s expected %d", actual, s, answer)
		}
	}
}
