package client

import (
	"log"

	"baseLearn/16.分布式爬搭建/crawler/engine"
	"baseLearn/16.分布式爬搭建/distributed/config"
	"baseLearn/16.分布式爬搭建/distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			// Call RPC to Save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return out, nil
}
