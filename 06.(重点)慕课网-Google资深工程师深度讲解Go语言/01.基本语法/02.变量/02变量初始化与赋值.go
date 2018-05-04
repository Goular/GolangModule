package main

import "fmt"

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func main() {
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
