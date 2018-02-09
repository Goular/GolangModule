package main

import "fmt"

func main() {
	/* 创建 map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	//删除元素
	delete(countryCapitalMap,"France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("删除元素后 map")
	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
}
