package main

import (
	"regexp"
	"fmt"
)

//第一匹配和最长匹配
func main() {
	b := []byte("abc1def1")
	pattern := `abc1|abc1def1`
	reg1 := regexp.MustCompile(pattern)// 第一匹配
	reg2 := regexp.MustCompilePOSIX(pattern)// 最长匹配
	fmt.Printf("%s\n",reg1.Find(b))
	fmt.Printf("%s\n",reg2.Find(b))
}
