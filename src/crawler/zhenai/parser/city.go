package parser

import (
	"crawler/engine"
	"crawler_distributed/config"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
var (
	profileRe = regexp.MustCompile(cityRe)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`) //不能随便添加换行符。
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents,-1)

	result :=engine.ParseResult{}
	for _, m:= range matches{
		url := string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items,"User " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url:url,
			//ParserFunc: ProfileParser(name),
			Parser:NewProfileParser(name),
		})

	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches{
		result.Requests = append(result.Requests,
			engine.Request{
				Url:string(m[1]),
				//ParserFunc:ParseCity,
				Parser:engine.NewFuncParser(ParseCity,config.ParseCity),
			})
	}

	return result

}