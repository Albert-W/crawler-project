package fetcher

import (
	"bufio"
	"crawler_distributed/config"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(
	time.Second / config.Qps)
//var rateLimiter = time.Tick(10 * time.Millisecond)
func Fetch(url string) ([]byte, error)  {
	<-rateLimiter
	//log.Printf("Fetching url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error: status code", resp.StatusCode)
		return nil,
		fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	//不要peek
	bodyReader :=bufio.NewReader(resp.Body)

	//自动find encoding
	//e :=determinEncoding(resp.Body)
	e :=determinEncoding(bodyReader)


	//GBK的转换
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}
	
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	//bytes, e := bufio.NewReader(r).Peek(1024)
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		//panic(e)
		log.Printf("Encoding error: %v", e)
		return unicode.UTF8
	}
	e2, _, _ := charset.DetermineEncoding(bytes, "")
	return e2
}