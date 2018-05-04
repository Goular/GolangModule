package main

import "fmt"

func main() {
	//createArrays()

	// 方法1:定义一个数组，默认存在零值
	var array1 [5] int
	// 方法2: 使用:=定义赋值符号必须添加数组的内容方法，如果是数组是定义，必须添加数组长度，不然就变成切片而不是数组
	array2 := [3]int{1, 2, 3}
	// 方法3: 数组定义，有定义赋值符号可以不添加准确的数组长度，可以使用省略号来代替长度，编译器会自动将后面的内容长度变为数组的长度，
	// 但是必须添加内容长度或者省略符号，不然就会变成切片而不是数组
	array3 := [...]int{4, 8, 7, 4, 95}

	printArray(array1)
	printArray(array2)
	printArray(array3)
}

// 打印数组
func printArray(arrs [5]int) {
	for _, value := range arrs {
		fmt.Print(value, " ")
	}
	fmt.Println("")
}

// 定义数组的方法
func createArrays() {
	// 方法1:定义一个数组，默认存在零值
	var array1 [5] int
	// 方法2: 使用:=定义赋值符号必须添加数组的内容方法，如果是数组是定义，必须添加数组长度，不然就变成切片而不是数组
	array2 := [3]int{1, 2, 3}
	// 方法3: 数组定义，有定义赋值符号可以不添加准确的数组长度，可以使用省略号来代替长度，编译器会自动将后面的内容长度变为数组的长度，
	// 但是必须添加内容长度或者省略符号，不然就会变成切片而不是数组
	array3 := [...]int{4, 8, 7, 4, 95, 48, 263}
	// 二维数组
	var grid [4][5]int // 这个写法存在零值
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(grid)
}
