package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	//Get请求使用
	//resp, err := http.Get("http://www.baidu.com")
	//defer resp.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//POST请求使用
	resp, err:=http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=101"))
	defer resp.Body.Close()
	if err != nil{
		panic(err)
	}
	body,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
