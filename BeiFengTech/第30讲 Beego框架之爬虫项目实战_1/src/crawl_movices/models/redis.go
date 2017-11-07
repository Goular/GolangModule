package models

import (
	"github.com/go-redis/redis"
)

const (
	URL_QUEUE     = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

var (
	client redis.Client
)

func ConnectRedis(addr string, password string, db int) {
	client = *redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func PutinQueue(url string) error {
	return client.LPush(URL_QUEUE, []byte(url)).Err()
}

func PopfromQueue() string {
	res, err := client.RPop(URL_QUEUE).Result()
	if err != nil {
		panic(err)
	}
	return string(res)
}

func AddToSet(url string) error {
	return client.SAdd(URL_VISIT_SET, []byte(url)).Err()
}

func IsVisit(url string) bool {
	bIsVisit, err := client.SIsMember(URL_VISIT_SET, []byte(url)).Result()
	if err != nil {
		return false
	}
	return bIsVisit
}

//测试是否联通
func CheckConnection() (string, error) {
	return client.Ping().Result()
}
