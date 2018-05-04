package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	//由于空字符串打不出来，所以我们使用格式化让其输出双引号
	fmt.Printf("%d %q", a, s)
}

func main() {
	variableZeroValue()
}
