package main

import "fmt"

//接口是实现覆盖方法使用的

//接口
type Phone interface {
	call()
}

//结构体
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {

}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}


func main() {
	var phone Phone;

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
