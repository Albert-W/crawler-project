package parser

import (
	"crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//contents, e := ioutil.ReadFile("citylist_test.html")

	if e!=nil {
		panic(e)
	}
	//fmt.Printf("%s\n",contents)

	result := ParseCityList(contents, "")
	const resultSize = 470
	//expectedUrls :=[]string{
	//	"","","",
	//}
	//expectedUrls :=[]string{
	//	"","","",
	//}

	if len(result.Requests) != resultSize {
		t.Errorf("result sould have %d requests; but had %d",
			resultSize, len(result.Requests))
	}
	//if len(result.Items) != resultSize {
	//	t.Errorf("result sould have %d requests; but had %d",
	//		resultSize, len(result.Items))
	//}

}
