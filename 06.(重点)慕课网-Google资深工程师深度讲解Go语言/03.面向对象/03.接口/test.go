package main

import (
	"fmt"
)

type One interface {
	Get(url string)
}

type Two interface {
	Post(url string)
}

type BothOneTwo interface {
	One
	Two
	Put(url string)
}

type S1 struct {
}

func (s1 *S1) String() string {
	return "S1"
}

func (s1 *S1) Get(url string) {
	fmt.Println("GET")
}

func (s1 *S1) Post(url string) {
	fmt.Println("POST")
}

func (s1 *S1) Put(url string) {
	fmt.Println("Put")
}

func main() {
	var one One = &S1{}
	one.Get("")
	var both BothOneTwo = &S1{}
	both.Post("")
	both.Put("")
	fmt.Println(&S1{})
}
