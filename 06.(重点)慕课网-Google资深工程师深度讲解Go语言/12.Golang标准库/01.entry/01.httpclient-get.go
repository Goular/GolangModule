package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://www.imooc.com", nil,
	)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.23 Mobile Safari/537.36")

	client := http.Client{
		CheckRedirect: func(req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect : ", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
