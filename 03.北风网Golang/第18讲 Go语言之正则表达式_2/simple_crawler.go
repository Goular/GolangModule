package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
)

func main() {
	url := "https://movie.douban.com/subject/1291832/?from=subject-page"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	sHtml, _ := ioutil.ReadAll(resp.Body)
	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(string(sHtml),-1)
	fmt.Println(result[0][1]);

	reg1 := regexp.MustCompile(`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`)
	result1 := reg1.FindAllStringSubmatch(string(sHtml), -1)

	fmt.Println(result1[0][1])
}
