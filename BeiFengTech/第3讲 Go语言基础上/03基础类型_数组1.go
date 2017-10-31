package main

import "fmt"
import "reflect"

var y [10]int


func main() {
	x:=[]int{1,2,3,4,5,6}
	fmt.Println("%v",x)
	y[0] = 22
	y[1] = 23
	y[2] = 24
	y[3] = 25
	y[4] = 26
	fmt.Println("%v",y)

	//x[7] = 2//执行这个会报错，因为数组长度是固定的
	fmt.Println(reflect.TypeOf(x))
}
