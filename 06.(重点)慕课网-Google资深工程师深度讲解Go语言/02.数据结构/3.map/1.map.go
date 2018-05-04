package main

import "fmt"

func main() {
	// 方式一
	m1 := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	// 方式二
	m2 := make(map[string]string) //等于empty
	// 方式三
	var m3 map[string]int // 等于nil
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	// 遍历map
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 如果判断获取map的属性存在还是不存在
	fmt.Println("------ Getting Value ------")

	if courseName, ok := m1["course2"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key is not exist")
	}

	// map删除元素
	fmt.Println("删除map中的元素")
	if _, ok := m1["name"]; ok {
		delete(m1, "name")
	} else {
		fmt.Println("key is not exist")
	}
	fmt.Println(m1)
}
