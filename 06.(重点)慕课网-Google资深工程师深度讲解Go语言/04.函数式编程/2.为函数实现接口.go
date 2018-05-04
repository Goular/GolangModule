package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 函数也能进行接口的实现
func (g intGen) Read(p []byte) (n int, err error) {
	next := g() //由于这里会不断的读，所以需要进行隔断
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// todo:如果数值太小的时候，会出错，因为[]byte作为缓存起来了，不够大，不输出
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci2()
	printFileContents(f)
}
