package main

import "fmt"

type Person struct {
	name string
	age  int8
}

type Student struct {
	Person
	specialStr string
}

func main() {
	s1:=Student{Person{"四维电脑",28},"学术性"}
	fmt.Printf("%v",s1)
	fmt.Println(s1.name)
}
