package models

import (
	"github.com/astaxie/goredis"
)

const (
	URL_QUEUE     = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

var (
	client goredis.Client
)

func ConnectRedis(addr string, port string, database_number int, password string) {
	//添加访问的参数
	client.Addr = addr + ":" + port
	client.Db = database_number
	client.Password = password
}

//将数据放入到Redis的列表队列中
func PushInQueue(url string) {
	client.Lpush(URL_QUEUE, []byte(url))
}

//将数据从列表队列中弹出
func PopFromQueue(url string) ([]byte, error) {
	return client.Rpop(URL_QUEUE)
}

//将数据添加到SET中
func AddToSet(url string) {
	client.Sadd(URL_VISIT_SET, []byte(url))
}

//判断特定的URL是否进行了浏览
func IsVisit(url string) (bool, error) {
	return client.Sismember(URL_VISIT_SET, []byte(url))
}
