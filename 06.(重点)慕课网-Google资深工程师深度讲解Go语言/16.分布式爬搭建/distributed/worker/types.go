package worker

import (
	"fmt"

	"log"

	"baseLearn/16.分布式爬搭建/crawler/config"
	"baseLearn/16.分布式爬搭建/crawler/engine"
	"baseLearn/16.分布式爬搭建/crawler/zhenai/parser"
	"github.com/pkg/errors"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := descrializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("Error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

// 将参数转变为函数(因为使用jsonrpc是不能够直接传递函数体，只能自建结构保存方法名和请求参数来传递到server方法进行处理)
func descrializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("Invalid arg: %v", p.Args)
		}
	default:
		return nil, errors.New("Unknow parser name")
	}
}
