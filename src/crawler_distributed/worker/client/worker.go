package client

import (
	"crawler/engine"
	"crawler_distributed/config"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"fmt"
)
//返回函数出来
func CreateProcessor()(engine.Processor,error)  {
	//起一个worker
	client, err := rpcsupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil{
		return nil,err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		//先序列化
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err :=client.Call(config.CrawlServiceRpc,sReq,&sResult)

		if err != nil{
			return engine.ParseResult{}, err
		}
		//反序列化
		return worker.DeserializeResult(sResult), nil

	}, nil

}

//func CreateProcessor(clientChan chan *rpc.Client) engine.Processor{
//
//	return func(req engine.Request) (engine.ParseResult, error) {
//		sReq := worker.SerializeRequest(req)
//		var sResult worker.ParseResult
//		client := <-clientChan
//		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
//		if err!=nil{
//			return engine.ParseResult{},err
//		}
//		return worker.DesrializeResult(sResult),nil
//	}
//}
