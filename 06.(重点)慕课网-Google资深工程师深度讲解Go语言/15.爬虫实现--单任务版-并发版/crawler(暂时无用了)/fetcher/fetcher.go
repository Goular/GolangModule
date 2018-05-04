package fetcher

import (
	"bufio"
	"io/ioutil"
	"net/http"

	"fmt"

	"log"

	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

// 抓取网页的内容
func Fetch(url string) ([]byte, error) {
	// 下面这样写，是为了，如果接收到定时的内容，我们在进行网站的爬取，这样会比较好
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong Status Code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 编码内容转义，防止中文不能正常读入
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 由于Reader不能直接读出来，不然就会没有的了，所以使用bufio来缓存处理
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error : %v", err)
		// 如果还是找不到相关的编码格式，默认为UTF-8
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
