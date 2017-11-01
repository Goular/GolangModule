package main

import "fmt"

//func (a int) compare(b,c int)bool  {
//	return a<b
//}

type Point struct {
	pX float32
	pY float32
}

//由于是方法参数值传递，可以看成是在内存空间复制一份出来运行操作，但是改变内容不会到对象
//的根内容，改变的知识复制对象，所以读取时是不存在作用的
func (point Point) setXY(x, y float32) {
	point.pX = x
	point.pY = y
}

//这个是指针传递，所以会产生内存根保存的真实变化
func (point *Point) setXY2(x, y float32) {
	point.pX = x
	point.pY = y
}

type Integer struct {
	value int
}

//Go是没有this对象的，但是可以定义任意引用对象作为参数，运行到函数体内部，这样才能更直观，
//运行和编译更快
func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}

func main() {
	p := Point{1.23, 4.56}
	p.setXY(3.21, 7.89)
	fmt.Println(p)
	p.setXY2(7.77, 8.99)
	fmt.Println(p)

}
