package main

import (
	"fmt"
	"regexp"
)

//const text = "My email is ccmouse@gmail.com"
const text = `
My email is ccmouse@gmail.com
email2: abc@126.com
email3: def@163.com.cn
`

func main() {
	// 1.普通搜索
	// re, _ := regexp.Compile(`[0-9A-Za-z]+@[0-9A-Za-z.]+\.[0-9A-Za-z]+`)
	// 2. 部分提取 -子匹配
	re, _ := regexp.Compile(`([0-9A-Za-z]+)@([0-9A-Za-z]+)(\.[0-9A-Za-z]+)`)
	// 1.
	//match := re.FindString(text)
	// 1.
	//match := re.FindAllString(text, -1)
	// 2. 搜索子匹配 返回的是一个二维的数组
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
