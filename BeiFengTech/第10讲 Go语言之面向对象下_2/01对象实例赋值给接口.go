package main

import "fmt"

/*
只能是打接口实现赋值给小接口，反之会报错
 */

type Bird struct {
}

type Animal interface {
	Run()
}

type Animal2 interface {
	Run()
	Walk()
}

func (bird Bird) Run() {
	fmt.Println("小鸟在跑步")
}

func (bird Bird) Walk() {
	fmt.Println("小鸟在走路")
}

func main() {
	var animal Animal
	var animal2 Animal2

	bird := new(Bird)
	animal2 = bird

	//尝试过animal2=animal会引起IDE报错，所以就没有使用例子了
	animal = animal2

	animal.Run()
}
