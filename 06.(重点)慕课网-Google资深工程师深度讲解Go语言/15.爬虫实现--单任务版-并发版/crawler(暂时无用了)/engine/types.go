package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

// 方法作用 :创建一个空的parse，这样是为了安全的定义类型
func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
