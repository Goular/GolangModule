package main

import (
	"regexp"
	"fmt"
)

func main() {
	reg1 := regexp.MustCompile("\\w+")
	reg2 := regexp.MustCompilePOSIX(`^z.*l$`)

	result1 := reg1.FindAllString("zhangsanl", -1)
	result2 := reg2.FindAllString("zhangsanl", -1)

	fmt.Printf("%v\n",result1)
	fmt.Printf("%v\n",result2)

	reg3 := regexp.MustCompile(`^z(.*)l$`)
	result3 := reg3.FindAllString("zhangsanl", -1)
	fmt.Printf("%v\n",result3)

	reg4 := regexp.MustCompile(`^z(.{1})(.{1})(.*)1$`)
	result4 := reg4.FindAllString("zhangsan1", -1)

	// 返回第一个匹配到的结果及其分组内容（结果以 b 的切片形式返回）。
	// 返回值中的第 0 个元素是整个正则表达式的匹配结果，后续元素是各个分组的
	// 匹配内容，分组顺序按照“(”的出现次序而定。
	//Submatch捕获内容使用的
	result5 := reg4.FindAllStringSubmatch("zhangsan1", -1)
	fmt.Printf("%v\n",result4)
	fmt.Printf("%v\n",result5)
}
