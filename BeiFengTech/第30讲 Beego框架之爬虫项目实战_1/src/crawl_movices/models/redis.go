package models

import (
	"github.com/astaxie/goredis"
)

const (
	URL_QUEUE = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

var (
	client goredis.Client
)

func ConnectRedis(addr string){
	client.Addr = addr
	//client.Password = "3071611103"
}

func PutinQueue(url string) error{
	return client.Lpush(URL_QUEUE, []byte(url))
}

func PopfromQueue() string{
	res,err := client.Rpop(URL_QUEUE)
	if err != nil{
		panic(err)
	}

	return string(res)
}

func AddToSet(url string){
	client.Sadd(URL_VISIT_SET, []byte(url))
}

func IsVisit(url string) bool{
	bIsVisit, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil{
		return false
	}

	return bIsVisit
}

