package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	// return
	panic("出现错误!")
	fmt.Println(4)

}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("Printed Too Many")
		}
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 写文件 -- 资源管理
func writeFile(fileName string) {
	// 创建一个文件资源
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 使用Writer来包装一下文件用于读取
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()

	tryDefer2()

	// 进行写文件
	//writeFile("fib.txt")
}
