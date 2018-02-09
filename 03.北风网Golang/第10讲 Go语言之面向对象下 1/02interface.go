package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}

type Bird struct {
}

func (bird Bird) Fly() {
	fmt.Println("Bird is flying!!!!")
}

func (bird Bird) Run() {
	fmt.Println("Bird is running!!!!")
}

func main() {
	//接口使用标准写法
	var animal1 Animal
	bird := new(Bird)
	animal1 = bird //将实例赋给接口作为其实现
	animal1.Fly()
	animal1.Run()
}
