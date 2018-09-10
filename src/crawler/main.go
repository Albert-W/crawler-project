package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	PaserFunc: parser.ParseCityList,
	//})
	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler:&scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker, //单机版直接调worker
	}
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	PaserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		//Url:       "http://www.zhenai.com/zhenghun/shanghai",
		Url: "http://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCity,
		Parser: engine.NewFuncParser(
			parser.ParseCityList, "ParseCityList"),
	})

	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//}
	////自动find encoding
	//e :=determinEncoding(resp.Body)
	//
	//
	////GBK的转换
	////utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Printf("%s\n", all)
	//printCityList(all)

}

//func printCityList(contents []byte)  {
//	exp := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"
//			[^>]*>([^<]+)</a>`
//	re := regexp.MustCompile(exp)
//	matchs := re.FindAllSubmatch(contents,-1)
//	for _, m:= range matchs{
//		fmt.Printf("City: %s, URL %s\n",m[2],m[1])
//	}
//	fmt.Printf("Count: %d\n",len(matchs))
//}
