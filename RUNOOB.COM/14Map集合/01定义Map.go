package main

import "fmt"

func main() {
	var countryMap map[string]string
	//创建集合
	countryMap = make(map[string]string)
	//插入相关的key-value值
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"

	//使用key 输出map的值
	for key, value := range countryMap {
		fmt.Println(key, value)
	}

	//查看元素在集合中是否存在
	key1,value1 := countryMap["USA"]
	if(value1){
		fmt.Println("Capital of United States is", key1)
	}else{
		fmt.Println("Capital of United States is not present")
	}
}
