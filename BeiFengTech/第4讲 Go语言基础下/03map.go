package main

import "fmt"

func main() {

	//下面呢那个是错误示范
	//var x map[string]float32
	//x["zhangsan"] = 91.65
	//fmt.Printf("%v",x)

	//正确使用map的例子
	x1 := make(map[string]float32)
	x1["Jack"] = 120.68
	x1["Tom"] = 125.98
	fmt.Printf("%v",x1)
}
