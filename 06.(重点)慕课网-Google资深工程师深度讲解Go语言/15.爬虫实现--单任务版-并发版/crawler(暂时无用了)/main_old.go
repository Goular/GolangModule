package main

import (
	"fmt"
	"net/http"

	"io/ioutil"

	"bufio"
	"io"

	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		// 暂不处理，以后再说
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error : Status code", resp.StatusCode)
		return
	}

	// 这里我们没有告诉他我们的是gbk，让他去猜和处理，最后返回一个encoding.Encoding，即编码解析器
	e := determineEncoding(resp.Body)

	// all为获取来的数据，由于返回的meta为GBK,所以会乱码
	// all, err := ioutil.ReadAll(resp.Body)
	// 需要使用golang/x/text库，将gbk转码为UTF-8
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", all) //打印完整的HTML数据
	printCityList(all)
}

// 处理返回的城市列表数据 ,城市列表解析器
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9A-Za-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//for _, subMatch := range m {
		// fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
		//}
		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
		fmt.Println()
	}
	fmt.Println(len(matches))
}

func determineEncoding(r io.Reader) encoding.Encoding {
	// 由于Reader不能直接读出来，不然就会没有的了，所以使用bufio来缓存处理
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
