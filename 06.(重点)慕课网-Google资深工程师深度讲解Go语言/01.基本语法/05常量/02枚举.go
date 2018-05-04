package main

import "fmt"

func enums() {
	const (
		cpp    = iota
		_
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
