package worker

import (
	"crawler/engine"
	"crawler_distributed/config"
	"crawler/zhenai/parser"
	"github.com/pkg/errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string  //函数名字
	Args interface{} //函数参数
	//{"ParseCityList",nil},{"ProfileParser",UserName(安静的雪)}
}



//可以在网上传递
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}
//进行序列化

//engine 下的Request
//type Request struct {
//	Url string
//	//ParserFunc ParserFunc
//	Parser Parser
//}
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	log.Printf("SerializeRequest "+name)
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}
//进行转换
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
//反序列化
func DeserializeRequest(r Request) (engine.Request, error) {
	parser1, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{},err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser1,
	},nil
}

func DeserializeResult(r ParseResult) engine.ParseResult{

	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err!=nil{
			log.Printf("erro deserilizing request:%v",err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	//还可以用map 维护
	switch p.Name {
	case config.ParseCityList:
		//log.Print("ParseCityList works well")
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), nil
	case config.ParseCity:
		//log.Print("ParseCity works well")
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	//case config.ParseProfile:
	case config.ParseProfile:
		log.Printf("ParseProfile")
		if userName, ok := p.Args.(string); ok {
			log.Print(userName)
			return parser.NewProfileParser(userName), nil
		} else {
			log.Print("ParseProfile falls")
			return nil, fmt.Errorf("invalid arg:%v", p.Args)
		}

	default:
		//log.Print(p.Name)//很关键的Debug
		return nil, errors.New("unknown parser name")

	}
}
