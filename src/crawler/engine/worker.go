package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker( r Request) (ParseResult,error) {
	//log.Printf("Fetching %s\n", r.Url)
	body, e := fetcher.Fetch(r.Url)
	if e != nil{
		log.Printf("Fetcher: error fetching url %s: %v",
			r.Url, e)
		return ParseResult{},e
	}

	//return r.ParserFunc(body,r.Url), nil
	return r.Parser.Parse(body,r.Url),nil
}