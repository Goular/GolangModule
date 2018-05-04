package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	// 读取文件的操作
	const fileName = "F://abc.txt"
	// 读取文件
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
