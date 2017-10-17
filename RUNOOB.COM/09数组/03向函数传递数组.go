package main

import "fmt"

/*使用的形参并为设定数组大小。*/
func main() {
	//定义数组的长度
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, 5);
	/* 输出返回的平均值 */
	fmt.Printf( "平均值为: %f ", avg );
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)

	return avg;
}
