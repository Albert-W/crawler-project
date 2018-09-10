package parser

import (
	"crawler/engine"
	"crawler_distributed/config"
	"regexp"
)
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"
			[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	//exp := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"
	//		[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents,-1)

	result :=engine.ParseResult{}
	//limit := 5
	for _, m:= range matchs{
		//result.Items = append(result.Items,"City "+string( m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			//ParserFunc: ParseCity,
			Parser:engine.NewFuncParser(ParseCity,config.ParseCity),
		})
		//fmt.Printf("City: %s, URL %s\n",m[2],m[1])
		//limit--
		//if limit ==0 {
		//	break
		//}

	}
	//fmt.Printf("Count: %d\n",len(matchs))
	return result
}
