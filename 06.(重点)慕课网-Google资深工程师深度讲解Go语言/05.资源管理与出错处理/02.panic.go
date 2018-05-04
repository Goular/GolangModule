package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-errors/errors"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 写文件 -- 资源管理
func writeFile2(fileName string) {
	// 创建一个文件资源
	file, err := os.OpenFile(
		// os.O_EXCL used with O_CREATE, file must not exist.
		fileName, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("This is a Custom error")
	if err != nil {
		//fmt.Println(err.Error())
		//return
		// panic(err)

		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err,
			)
		}
		return
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
	writeFile2("fib.txt")
}
