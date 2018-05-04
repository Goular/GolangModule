package main

import (
	"fmt"

	"time"

	"baseLearn/03.面向对象/03.接口/mock"
	"baseLearn/03.面向对象/03.接口/real"
)

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

// 接口组合
type RetrieverPoster interface {
	Retriever
	Poster
	Connect(host string)
}

const url = "http://www.imooc.com"

// 组合接口的使用
func session(s RetrieverPoster) {
	s.Post(url, map[string]string{
		"contents ": "another faked imooc.com",
	})
	s.Get(url)
}

func main() {
	// 定义接口
	var r Retriever
	// 为接口进行赋值
	r = &mock.Retriever{"This is a fake imooc.com"}
	inspect(r)
	fmt.Printf("%T %v\n", r, r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	// 断言
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
