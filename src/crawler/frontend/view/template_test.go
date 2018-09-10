package view

import (
	"crawler/engine"
	"crawler/frontend/view/model"
	common "crawler/model"
	"os"
	"testing"
)

func TestS(t *testing.T)  {
	view := CreateSearchResultView("view/template.test.html")

	//template:= template.Must(
	//template.ParseFiles("view/template.html"))

	out, err := os.Create("view/template.test.html")

	page := model.SearchResult{}

	page.Hits = 123

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: common.Profile{
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

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items,item)
	}

	//err = template.Execute(out, page)
	err = view.Render(out,page)
	if err != nil{
		panic(err)
	}
}