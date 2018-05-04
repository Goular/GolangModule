package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// for 的使用
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(72387885),
	)
	printFile("F://abc.txt")
	//forever()
	s := `
	Hello World.
	123456
	popppp
`
	printFileContents(strings.NewReader(s))
}

func convertToBin(v int) string {
	result := ""
	for ; v > 0; v /= 2 {
		result = strconv.Itoa(v%2) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}
