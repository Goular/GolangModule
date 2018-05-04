package parse

import (
	"regexp"

	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/engine"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

// 解析城市列表的内容
func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// 返回会员的名字
		//result.Items = append(result.Items, "User "+string(m[2]))
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
