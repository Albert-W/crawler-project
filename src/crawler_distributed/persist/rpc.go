package persist

import (
	"crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}
//调用Save服务
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.",item)
	if err == nil{
		*result="ok"
	}else {
		log.Printf("Error saving item %v:%v",item,err)
	}
	return err
}
