package engine

//带两个函数的接口
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParserFunc func(contents []byte, url string) ParseResult
type Request struct {
	Url string
	//ParserFunc ParserFunc
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string //存储时去重。
	Url     string
	Type    string //存储的配置
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}
// 函数类型的Parser
type FuncParser struct {
	parser ParserFunc //对应解析函数
	name   string //函数名
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}




//func CreateFuncParser(
//	p ParserFunc, name string) *FuncParser {
//	return &FuncParser{
//		parser: p,
//		name:   name,
//	}
//}

//func NilParser([]byte) ParseResult  {
//	return ParseResult{}
//}

//type SerializedParser struct {
//	Name string
//	Args interface{}
//}

//{"ParseCitylist", nil}
//{"ProfileParser", userName}
