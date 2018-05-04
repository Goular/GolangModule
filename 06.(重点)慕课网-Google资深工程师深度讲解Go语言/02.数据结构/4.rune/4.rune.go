package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes 我爱慕课网!"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	// 打印rune长度
	fmt.Println("Rune Count :", utf8.RuneCountInString(s))
	// 解码byte
	bytes := []byte(s)
	for len(bytes) > 0 {
		// 每一次解完，获取最后的位置用于切片截取
		ch, size := utf8.DecodeRune(bytes) // 根据rune来解析
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	// 直接使用rune来解码字符串
	for index, char := range []rune(s) {
		fmt.Printf("(%d , %c )", index, char)
	}
	fmt.Println()

}
