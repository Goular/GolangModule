package main

import "fmt"

func main() {
	x := make([]int,3,5)
	x = append(x, 2,4,5,6)

	//展示内容
	fmt.Printf("%v\n",x)
	//代表现在slice中元素的个数
	fmt.Printf("%v\n",len(x))
	//代表现在slice底层的分配的空间是多大 当append加入的元素个数超过此值时 底层会自动扩充
	fmt.Printf("%v\n",cap(x))

}
