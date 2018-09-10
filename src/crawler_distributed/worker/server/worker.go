package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

//命令行库，可以读取加入的参数
// go run worker.go --port=9000
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	//log.Fatal(rpcsupport.ServeRpc(
	//	fmt.Sprintf(":%d",config.WorkerPort0),
	//	worker.CrawlService{}))
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
