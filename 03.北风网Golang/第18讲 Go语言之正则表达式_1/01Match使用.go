package main

import (
	"regexp"
	"fmt"
)

func main() {
	//判断是否在搜索字符串中找到匹配pattern所匹配的字符串
	//pattern := 	`(((abc.)def.)ghi)`
	//src := "abc-def-ghi abc+def+ghi"
	//result1 ,_:= regexp.Match(pattern,[]byte(src))
	//fmt.Println(result1)
	//fmt.Println(regexp.QuoteMeta(pattern))

	pattern := 	`(((abc.)def.)ghi)`
	src := "abc-def-ghi abc+def+ghi"
	result1 ,_:= regexp.MatchString(pattern,src)
	fmt.Println(result1)
	fmt.Println(regexp.QuoteMeta(pattern))
}
