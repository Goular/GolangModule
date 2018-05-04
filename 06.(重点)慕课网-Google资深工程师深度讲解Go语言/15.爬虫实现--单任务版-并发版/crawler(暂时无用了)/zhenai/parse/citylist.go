package parse

import (
	"regexp"

	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9A-Za-z]+)"[^>]*>([^<]+)</a>`

// 解析城市列表的内容
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//limit := 10
	for _, m := range matches {
		// 返回城市的名字
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		// 减少页数，使尽快出Profile数据
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
