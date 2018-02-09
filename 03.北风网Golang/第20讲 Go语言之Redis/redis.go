package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "IP地址:6379",
		Password: "密码",
		DB:       0,
	})
	CheckConnection(client)

	SimpleUse(client)
}

//简单使用
func SimpleUse(client *redis.Client) {
	//字符串使用
	err := client.Set("test", "Hello JiaGongWu", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	//错误处理
	val2, err := client.Get("keys").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	//Map的使用
	f := make(map[string]interface{})
	f["name"] = "JiaGongWu"
	f["age"] = 12
	f["sex"] = "male2"

	fmt.Println(f)

	client.HMSet("JiaGongWu",f).Err()
	//if err != nil {
	//	panic(err2)
	//}
	val3,err := client.HMGet("JiaGongWu","name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val3)
}

//测试是否联通
func CheckConnection(client *redis.Client) {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
