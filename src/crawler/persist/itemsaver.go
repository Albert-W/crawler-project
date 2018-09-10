package persist

import (
	"crawler/engine"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, e := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if e != nil {
		return nil,e
	}

	out := make(chan engine.Item)
	go func() {
		itemCount :=0
		for {
			//用于存储item
			item :=<- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil{
				log.Printf("Item saver: error saving item %v: %v",
					item, err)

			}
		}
	}()
	return out,nil

}
func Save(client *elastic.Client ,index string , item engine.Item)  error {


	if item.Type == ""{
		return errors.New("Must supply type")
	}

	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != ""{
		indexService.Id(item.Id)
	}

	_ , err := indexService.
		Do(context.Background())
	if err!=nil{
		return err
	}
	//fmt.Printf("%+v",resp)
	return nil
}