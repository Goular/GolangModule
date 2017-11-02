package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	//小写会隐藏外部访问
	//name string
	//age  int
	Name string `json:"StudentName"` //为json对象的key起别名
	Age int
}

func main() {
	//对数组进行json编码
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	//对Map进行JSON编码
	m := make(map[string]string);
	m["GuangZhou"] = "羊城"
	m["ShangHai"] = "申城"
	m["BeiJing"] = "京城"
	s2, err2 := json.Marshal(m);
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(s2))

	//对对象进行json编码
	o := new(Student)
	o.Name = "Goular"
	o.Age = 25
	s3, err3 := json.Marshal(o)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(s3))

	//对s3进行解码,对象会返回的是map
	var v4 interface{}
	json.Unmarshal([]byte(s3),&v4)
	fmt.Println(v4)

	//对s3进行解码,map会返回的是map
	var v5 interface{}
	json.Unmarshal([]byte(s2),&v5)
	fmt.Println(v5)
}
