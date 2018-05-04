package engine

import (
	"log"

	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/fetcher"
)

type SimpleEngine struct {
}

// seeds 为种子请求集合
func (engine SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		// 回去的是请求的地址
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v ", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
