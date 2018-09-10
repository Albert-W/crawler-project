package persist

import (
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)
//测试了先存后取。
func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "....",
		Type: "zhenai",
		Id:   "1089",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "",
			Occupation: "",
			Marriage:   "",
			House:      "",
			Hokou:      "",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	// TODO: Try to start up elastic search here using docker go client
	client, e := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}

	const index = "dating_test"
	// Save expected item
	err := Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s",resp.Source)

	var actual engine.Item
	err = json.Unmarshal(
		*resp.Source, &actual)
	//if err != nil {
	//	panic(err)
	//}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//verify result
	if actual != expected {
		t.Errorf("got %v, expected %v",
			actual,expected)
	}
}
