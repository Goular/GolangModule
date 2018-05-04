package main

import (
	"fmt"

	"baseLearn/03.面向对象/02.queue"
)

func main() {
	q := _2_queue.Queue{1}
	var Q1 *_2_queue.Queue = &q
	fmt.Println(Q1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(&q)
}
