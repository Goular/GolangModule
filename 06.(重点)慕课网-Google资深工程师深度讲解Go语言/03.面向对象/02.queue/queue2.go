package queue

// 通用类型版队列
// 通用版的缺陷是类型需要在运行的时候才能动态判断
type Queue2 []interface{}

func (q *Queue2) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue2) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue2) IsEmpty() bool {
	return len(*q) == 0
}
