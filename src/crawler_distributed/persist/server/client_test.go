package main

import (
	"testing"
	"crawler_distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"time"
	"crawler_distributed/config"
)

func TestItemSaver(t *testing.T) {
	// start ItemSaverServer
	// start ItemSaverClient
	// Call save

	const host = ":1234"
	//启动服务器
	go serveRpc(host, "test1")
	time.Sleep(10*time.Second)  //时间长一些，保证go执行完
	//启动客户端
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//调用相应接口
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牧羊座",
			Occupation: "人事/行政",
			Marriage:    "离异",
			House:      "已够房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok"{
		t.Errorf("result:%s; err: %s",result,err)
	}
}
