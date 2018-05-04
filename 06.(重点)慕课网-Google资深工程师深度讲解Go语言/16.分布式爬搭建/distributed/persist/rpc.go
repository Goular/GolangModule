package persist

import (
	"baseLearn/16.分布式爬搭建/crawler/engine"

	"log"

	"baseLearn/16.分布式爬搭建/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return err
}
