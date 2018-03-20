package main
// 模拟生成相关数据项
import (
	"flag"
	"fmt"
	"strings"
	"strconv"
	"net/url"
	"math/rand"
	"time"
	"os"
)

var uaList = []string{
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
	"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
	"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; TencentTraveler 4.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SE 2.X MetaSr 1.0; SE 2.X MetaSr 1.0; .NET CLR 2.0.50727; SE 2.X MetaSr 1.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Avant Browser)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",
}

type resource struct {
	url    string
	target string
	start  int
	end    int
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{// 首页
		url: "http://localhost:8888/",
		target: "",
		start: 0,
		end: 0,
	}
	r2 := resource{// 列表页
		url: "http://localhost:8888/list/{$id}.html",
		target: "{$id}",
		start: 1,
		end: 21,
	}
	r3 := resource{//
		url: "http://localhost:8888/movie/{$id}.html",
		target: "{$id}",
		start: 1,
		end: 12924,
	}
	res = append(append(append(res, r1), r2), r3)
	return res
}

func buildUrl(res []resource) []string {
	var list []string
	for _, item := range res {
		if len(item.target) == 0 {
			list = append(list, item.url)
		} else {
			for i := item.start; i <= item.end; i++ {
				urlStr := strings.Replace(item.url, item.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func main() {
	total := flag.Int("total", 100, "How many rows by created")
	filePath := flag.String("filePath", "E:\\nginx\\logs\\dig.log", "Log File Path.")
	flag.Parse()
	//fmt.Println(*total, *filePath)
	// 需要构造出真实的网站的URL的集合
	res := ruleResource()
	list := buildUrl(res)

	//fmt.Println(list)
	//return

	logStr := ""
	//fmt.Println(list)
	// 按照要求，生成$total行日志内容，数据来源源自上面的集合
	for i := 0; i <= *total; i++ {
		currentUrl := list[randInt(0, len(list)-1)]
		referUrl := list[randInt(0, len(list)-1)]
		ua := uaList[randInt(0, len(uaList)-1)]
		logStr += makeLog(currentUrl, referUrl, ua) + "\n"
		time.Sleep(2 * time.Millisecond)
	}
	fd, _ := os.OpenFile(*filePath, os.O_RDWR|os.O_APPEND, 0644)
	fd.Write([]byte(logStr))
	fd.Close()
	fmt.Println("done.\n")
}
func makeLog(current, refer, ua string) string {
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramsStr := u.Encode()
	logTemplate := "127.0.0.1 - - [16/Mar/2018:15:03:13 +0800] \"OPTIONS /dig?{$paramsStr} HTTP/1.1\" 200 43 \"-\" \"{$ua}\" \"-\""
	log := strings.Replace(logTemplate, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)
	return log
}
